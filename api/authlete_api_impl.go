//
// Copyright (C) 2019-2021 Authlete, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the
// License.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/authlete/authlete-go/conf"
	"github.com/authlete/authlete-go/dto"
)

type impl struct {
	baseUrl               string
	serviceOwnerApiKey    string
	serviceOwnerApiSecret string
	serviceApiKey         string
	serviceApiSecret      string
	settings              *Settings
}

func (self *impl) init(configuration conf.AuthleteConfiguration) {
	self.baseUrl = configuration.GetBaseUrl()
	self.serviceOwnerApiKey = configuration.GetServiceOwnerApiKey()
	self.serviceOwnerApiSecret = configuration.GetServiceOwnerApiSecret()
	self.serviceApiKey = configuration.GetServiceApiKey()
	self.serviceApiSecret = configuration.GetServiceApiSecret()
	self.settings = &Settings{}
}

func (self *impl) callApi(
	method string, apiKey string, apiSecret string, path string,
	queryParams map[string]string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	// Prepare a request to the Authlete API.
	req := self.buildRequest(method, apiKey, apiSecret, path, queryParams, requestBody)

	// HTTP Client
	client := self.prepareClient()

	// Call the Authlete API.
	res, err := client.Do(req)
	if err != nil {
		msg := fmt.Sprintf(`API call to '%s' failed`, path)
		return &AuthleteError{err, msg, req, res}
	}

	// HTTP Status Code of the API response.
	sc := res.StatusCode
	if sc < 200 || 300 <= sc {
		msg := fmt.Sprintf(`'%s' API returned %d`, path, sc)
		return &AuthleteError{nil, msg, req, res}
	}

	// Fill the object with the content of the response body.
	err = self.deserializeResponseBody(res, responseContainer)
	if err != nil {
		msg := fmt.Sprintf(`Failed to deserialize the response from '%s' API`, path)
		return &AuthleteError{err, msg, req, res}
	}

	// No error
	return nil
}

func (self *impl) buildRequest(
	method string, apiKey string, apiSecret string, path string,
	queryParams map[string]string, requestBody interface{}) *http.Request {
	req := http.Request{}

	req.Method = method
	req.Header = self.buildRequestHeader()
	req.URL = self.buildRequestUrl(path, queryParams)
	req.Body = self.buildRequestBody(requestBody)
	req.SetBasicAuth(apiKey, apiSecret)

	return &req
}

func (self *impl) buildRequestHeader() http.Header {
	return map[string][]string{
		`Accept`:       {`application/json`},
		`Content-Type`: {`application/json;charset=UTF-8`},
	}
}

func (self *impl) buildRequestUrl(path string, queryParams map[string]string) *url.URL {
	var builder strings.Builder

	builder.WriteString(self.baseUrl)
	builder.WriteString(path)

	if queryParams != nil {
		delimiter := `?`
		for key, value := range queryParams {
			builder.WriteString(delimiter)
			builder.WriteString(url.QueryEscape(key))
			builder.WriteString(`=`)
			builder.WriteString(url.QueryEscape(value))
			delimiter = `&`
		}
	}

	rawurl := builder.String()
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Panicf(`Failed to build a URL from '%v'`, rawurl)
	}

	return u
}

func (self *impl) buildRequestBody(requestBody interface{}) io.ReadCloser {
	if requestBody == nil {
		return nil
	}

	str, ok := requestBody.(string)
	if ok {
		// Convert string to io.ReadCloser
		return ioutil.NopCloser(strings.NewReader(str))
	}

	// Convert the object into []byte.
	byt, err := json.Marshal(requestBody)
	if err != nil {
		log.Panic(`Failed to marshal the object`)
	}

	// Convert []byte to io.ReadCloser
	return ioutil.NopCloser(bytes.NewReader(byt))
}

func (self *impl) prepareClient() *http.Client {
	return &http.Client{
		Timeout: self.settings.Timeout,
	}
}

func (self *impl) deserializeResponseBody(
	res *http.Response, responseContainer interface{}) error {
	if responseContainer == nil {
		return nil
	}

	defer res.Body.Close()

	// Read the body of the response.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	sb, ok := responseContainer.(*strings.Builder)
	if ok {
		// Convert []byte to string.
		sb.WriteString(string(body))
		return nil
	}

	// Convert the JSON (body) to an instance.
	err = json.Unmarshal(body, responseContainer)
	if err != nil {
		return err
	}

	return nil
}

func (self *impl) callGetApi(
	apiKey string, apiSecret string, path string, queryParams map[string]string,
	responseContainer interface{}) *AuthleteError {
	return self.callApi(
		http.MethodGet, apiKey, apiSecret,
		path, queryParams, nil, responseContainer)
}

func (self *impl) callPostApi(
	apiKey string, apiSecret string, path string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	return self.callApi(
		http.MethodPost, apiKey, apiSecret,
		path, nil, requestBody, responseContainer)
}

func (self *impl) callDeleteApi(
	apiKey string, apiSecret string, path string) *AuthleteError {
	return self.callApi(
		http.MethodDelete, apiKey, apiSecret, path, nil, nil, nil)
}

func (self *impl) callServiceGetApi(
	path string, queryParams map[string]string,
	responseContainer interface{}) *AuthleteError {
	return self.callGetApi(
		self.serviceApiKey, self.serviceApiSecret,
		path, queryParams, responseContainer)
}

func (self *impl) callServicePostApi(
	path string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	return self.callPostApi(
		self.serviceApiKey, self.serviceApiSecret,
		path, requestBody, responseContainer)
}

func (self *impl) callServiceDeleteApi(path string) *AuthleteError {
	return self.callDeleteApi(
		self.serviceApiKey, self.serviceApiSecret, path)
}

func (self *impl) callServiceOwnerGetApi(
	path string, queryParams map[string]string,
	responseContainer interface{}) *AuthleteError {
	return self.callGetApi(
		self.serviceOwnerApiKey, self.serviceOwnerApiSecret,
		path, queryParams, responseContainer)
}

func (self *impl) callServiceOwnerPostApi(
	path string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	return self.callPostApi(
		self.serviceOwnerApiKey, self.serviceOwnerApiSecret,
		path, requestBody, responseContainer)
}

func (self *impl) callServiceOwnerDeleteApi(path string) *AuthleteError {
	return self.callDeleteApi(
		self.serviceOwnerApiKey, self.serviceOwnerApiSecret, path)
}

func buildMap(args ...interface{}) map[string]string {
	// The implementation of this method assumes that the number of
	// the arguments is even and that keys and values appear alternately
	// (e.g. `key0`, `value0`, `key1`, `value1`, ...)

	m := map[string]string{}

	for i := 0; i < len(args); i += 2 {
		key := args[i].(string)
		val := args[i+1]

		if val != nil {
			m[key] = toString(val)
		}
	}

	return m
}

func toString(val interface{}) string {
	switch v := val.(type) {
	case bool:
		return strconv.FormatBool(v)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case string:
		return v
	default:
		return fmt.Sprintf(`%v`, v)
	}
}

func (self *impl) Settings() *Settings {
	return self.settings
}

func (self *impl) Authorization(
	request *dto.AuthorizationRequest) (res *dto.AuthorizationResponse, err *AuthleteError) {
	res = &dto.AuthorizationResponse{}
	err = self.callServicePostApi(`/api/auth/authorization`, request, res)
	return
}

func (self *impl) AuthorizationFail(
	request *dto.AuthorizationFailRequest) (res *dto.AuthorizationFailResponse, err *AuthleteError) {
	res = &dto.AuthorizationFailResponse{}
	err = self.callServicePostApi(`/api/auth/authorization/fail`, request, res)
	return
}

func (self *impl) AuthorizationIssue(
	request *dto.AuthorizationIssueRequest) (res *dto.AuthorizationIssueResponse, err *AuthleteError) {
	res = &dto.AuthorizationIssueResponse{}
	err = self.callServicePostApi(`/api/auth/authorization/issue`, request, res)
	return
}

func (self *impl) Token(
	request *dto.TokenRequest) (res *dto.TokenResponse, err *AuthleteError) {
	res = &dto.TokenResponse{}
	err = self.callServicePostApi(`/api/auth/token`, request, res)
	return
}

func (self *impl) TokenCreate(
	request *dto.TokenCreateRequest) (res *dto.TokenCreateResponse, err *AuthleteError) {
	res = &dto.TokenCreateResponse{}
	err = self.callServicePostApi(`/api/auth/token/create`, request, res)
	return
}

func (self *impl) TokenDelete(
	token string) (err *AuthleteError) {
	path := `/api/auth/token/delete/` + token
	err = self.callServiceDeleteApi(path)
	return
}

func (self *impl) TokenFail(
	request *dto.TokenFailRequest) (res *dto.TokenFailResponse, err *AuthleteError) {
	res = &dto.TokenFailResponse{}
	err = self.callServicePostApi(`/api/auth/token/fail`, request, res)
	return
}

func (self *impl) TokenIssue(
	request *dto.TokenIssueRequest) (res *dto.TokenIssueResponse, err *AuthleteError) {
	res = &dto.TokenIssueResponse{}
	err = self.callServicePostApi(`/api/auth/token/issue`, request, res)
	return
}

func (self *impl) TokenUpdate(
	request *dto.TokenUpdateRequest) (res *dto.TokenUpdateResponse, err *AuthleteError) {
	res = &dto.TokenUpdateResponse{}
	err = self.callServicePostApi(`/api/auth/token/update`, request, res)
	return
}

func (self *impl) GetTokenList(
	clientIdentifier string, subject string, start uint32, end uint32) (res *dto.TokenListResponse, err *AuthleteError) {
	params := buildMap(
		`clientIdentifier`, clientIdentifier,
		`subject`, subject, `start`, start, `end`, end)

	res = &dto.TokenListResponse{}
	err = self.callServiceGetApi(`/api/auth/token/get/list`, params, res)
	return
}

func (self *impl) Revocation(
	request *dto.RevocationRequest) (res *dto.RevocationResponse, err *AuthleteError) {
	res = &dto.RevocationResponse{}
	err = self.callServicePostApi(`/api/auth/revocation`, request, res)
	return
}

func (self *impl) UserInfo(
	request *dto.UserInfoRequest) (res *dto.UserInfoResponse, err *AuthleteError) {
	res = &dto.UserInfoResponse{}
	err = self.callServicePostApi(`/api/auth/userinfo`, request, res)
	return
}

func (self *impl) UserInfoIssue(
	request *dto.UserInfoIssueRequest) (res *dto.UserInfoIssueResponse, err *AuthleteError) {
	res = &dto.UserInfoIssueResponse{}
	err = self.callServicePostApi(`/api/auth/userinfo/issue`, request, res)
	return
}

func (self *impl) Introspection(
	request *dto.IntrospectionRequest) (res *dto.IntrospectionResponse, err *AuthleteError) {
	res = &dto.IntrospectionResponse{}
	err = self.callServicePostApi(`/api/auth/introspection`, request, res)
	return
}

func (self *impl) StandardIntrospection(
	request *dto.StandardIntrospectionRequest) (res *dto.StandardIntrospectionResponse, err *AuthleteError) {
	res = &dto.StandardIntrospectionResponse{}
	err = self.callServicePostApi(`/api/auth/introspection/standard`, request, res)
	return
}

func (self *impl) CreateService(
	request *dto.Service) (res *dto.Service, err *AuthleteError) {
	res = &dto.Service{}
	err = self.callServiceOwnerPostApi(`/api/service/create`, request, res)
	return
}

func (self *impl) DeleteService(
	apiKey interface{}) (err *AuthleteError) {
	path := `/api/service/delete/` + toString(apiKey)
	err = self.callServiceOwnerDeleteApi(path)
	return
}

func (self *impl) GetService(
	apiKey interface{}) (res *dto.Service, err *AuthleteError) {
	path := `/api/service/get/` + toString(apiKey)
	res = &dto.Service{}
	err = self.callServiceOwnerGetApi(path, nil, res)
	return
}

func (self *impl) GetServiceList(
	start uint32, end uint32) (res *dto.ServiceListResponse, err *AuthleteError) {
	params := buildMap(`start`, start, `end`, end)

	res = &dto.ServiceListResponse{}
	err = self.callServiceOwnerGetApi(`/api/service/get/list`, params, res)
	return
}

func (self *impl) UpdateService(
	service *dto.Service) (res *dto.Service, err *AuthleteError) {
	path := `/api/service/update/` + toString(service.ApiKey)
	res = &dto.Service{}
	err = self.callServiceOwnerPostApi(path, service, res)
	return
}

func (self *impl) GetServiceJwks(
	pretty bool, includePrivateKeys bool) (res string, err *AuthleteError) {
	params := buildMap(`pretty`, pretty, `includePrivateKeys`, includePrivateKeys)

	obj := strings.Builder{}
	err = self.callServiceGetApi(`/api/service/jwks/get`, params, &obj)

	if err == nil {
		res = obj.String()
	}

	return
}

func (self *impl) GetServiceConfiguration(
	pretty bool) (res string, err *AuthleteError) {
	params := buildMap(`pretty`, pretty)

	obj := strings.Builder{}
	err = self.callServiceGetApi(`/api/service/configuration`, params, &obj)

	if err == nil {
		res = obj.String()
	}

	return
}

func (self *impl) CreateClient(
	client *dto.Client) (res *dto.Client, err *AuthleteError) {
	res = &dto.Client{}
	err = self.callServicePostApi(`/api/client/create`, client, res)
	return
}

func (self *impl) DynamicClientRegister(
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	res = &dto.ClientRegistrationResponse{}
	err = self.callServicePostApi(`/api/client/registration`, request, res)
	return
}

func (self *impl) DynamicClientGet(
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	res = &dto.ClientRegistrationResponse{}
	err = self.callServicePostApi(`/api/client/registration/get`, request, res)
	return
}

func (self *impl) DynamicClientUpdate(
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	res = &dto.ClientRegistrationResponse{}
	err = self.callServicePostApi(`/api/client/registration/update`, request, res)
	return
}

func (self *impl) DynamicClientDelete(
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	res = &dto.ClientRegistrationResponse{}
	err = self.callServicePostApi(`/api/client/registration/delete`, request, res)
	return
}

func (self *impl) DeleteClient(
	clientIdentifier interface{}) (err *AuthleteError) {
	path := `/api/client/delete/` + toString(clientIdentifier)
	err = self.callServiceDeleteApi(path)
	return
}

func (self *impl) GetClient(
	clientIdentifier interface{}) (res *dto.Client, err *AuthleteError) {
	path := `/api/client/get/` + toString(clientIdentifier)
	res = &dto.Client{}
	err = self.callServiceGetApi(path, nil, res)
	return
}

func (self *impl) GetClientList(
	developer string, start uint32, end uint32) (res *dto.ClientListResponse, err *AuthleteError) {
	params := buildMap(`developer`, developer, `start`, start, `end`, end)

	res = &dto.ClientListResponse{}
	err = self.callServiceGetApi(`/api/client/get/list`, params, res)
	return
}

func (self *impl) UpdateClient(
	client *dto.Client) (res *dto.Client, err *AuthleteError) {
	path := `/api/client/update/` + toString(client.ClientId)
	res = &dto.Client{}
	err = self.callServicePostApi(path, client, res)
	return
}

type requestableScopes struct {
	RequestableScopes []string `json:"requestableScopes,omitempty"`
}

func (self *impl) GetRequestableScopes(
	clientIdentifier interface{}) (res []string, err *AuthleteError) {
	path := `/api/client/extension/requestable_scopes/get/` + toString(clientIdentifier)

	obj := requestableScopes{}
	err = self.callServiceGetApi(path, nil, &obj)

	if err == nil {
		res = obj.RequestableScopes
	}

	return
}

func (self *impl) SetRequestableScopes(
	clientIdentifier interface{}, scopes []string) (res []string, err *AuthleteError) {
	input := requestableScopes{}
	input.RequestableScopes = scopes

	path := `/api/client/extension/requestable_scopes/update/` + toString(clientIdentifier)

	output := requestableScopes{}
	err = self.callServicePostApi(path, &input, &output)

	if err == nil {
		res = output.RequestableScopes
	}

	return
}

func (self *impl) DeleteRequestableScopes(
	clientIdentifier interface{}) (err *AuthleteError) {
	path := `/api/client/extension/requestable_scopes/delete/` + toString(clientIdentifier)
	err = self.callServiceDeleteApi(path)
	return
}

type grantedScopesRequest struct {
	Subject string `json:"subject,omitempty"`
}

func (self *impl) GetGrantedScopes(
	clientIdentifier interface{}, subject string) (res *dto.GrantedScopesGetResponse, err *AuthleteError) {
	req := grantedScopesRequest{}
	req.Subject = subject

	path := `/api/client/granted_scopes/get/` + toString(clientIdentifier)

	res = &dto.GrantedScopesGetResponse{}
	err = self.callServicePostApi(path, &req, res)
	return
}

func (self *impl) DeleteGrantedScopes(
	clientIdentifier interface{}, subject string) (err *AuthleteError) {
	path := `/api/client/granted_scopes/delete/` + toString(clientIdentifier)
	err = self.callServiceDeleteApi(path)
	return
}

func (self *impl) DeleteClientAuthorization(
	clientIdentifier interface{}, subject string) (err *AuthleteError) {
	req := dto.ClientAuthorizationDeleteRequest{}
	req.Subject = subject

	path := `/api/client/authorization/delete/` + toString(clientIdentifier)

	res := &dto.ApiResponse{}
	err = self.callServicePostApi(path, &req, res)
	return
}

func (self *impl) GetClientAuthorizationList(
	request *dto.ClientAuthorizationGetListRequest) (res *dto.AuthorizedClientListResponse, err *AuthleteError) {
	res = &dto.AuthorizedClientListResponse{}
	err = self.callServicePostApi(`/api/client/authorization/get/list`, request, res)
	return
}

func (self *impl) UpdateClientAuthorization(
	clientIdentifier interface{}, request *dto.ClientAuthorizationUpdateRequest) (err *AuthleteError) {
	path := `/api/client/authorization/update/` + toString(clientIdentifier)

	res := &dto.ApiResponse{}
	err = self.callServicePostApi(path, request, res)
	return
}

func (self *impl) RefreshClientSecret(
	clientIdentifier interface{}) (res *dto.ClientSecretRefreshResponse, err *AuthleteError) {
	path := `/api/client/secret/refresh/` + toString(clientIdentifier)

	res = &dto.ClientSecretRefreshResponse{}
	err = self.callServiceGetApi(path, nil, res)
	return
}

func (self *impl) UpdateClientSecret(
	clientIdentifier interface{}, clientSecret string) (res *dto.ClientSecretUpdateResponse, err *AuthleteError) {
	req := dto.ClientSecretUpdateRequest{}
	req.ClientSecret = clientSecret

	path := `/api/client/secret/update/` + toString(clientIdentifier)

	res = &dto.ClientSecretUpdateResponse{}
	err = self.callServicePostApi(path, &req, res)
	return
}

func (self *impl) VerifyJose(
	request *dto.JoseVerifyRequest) (res *dto.JoseVerifyResponse, err *AuthleteError) {
	res = &dto.JoseVerifyResponse{}
	err = self.callServicePostApi(`/api/jose/verify`, request, res)
	return
}

func (self *impl) BackchannelAuthentication(
	request *dto.BackchannelAuthenticationRequest) (res *dto.BackchannelAuthenticationResponse, err *AuthleteError) {
	res = &dto.BackchannelAuthenticationResponse{}
	err = self.callServicePostApi(`/api/backchannel/authentication`, request, res)
	return
}

func (self *impl) BackchannelAuthenticationIssue(
	request *dto.BackchannelAuthenticationIssueRequest) (res *dto.BackchannelAuthenticationIssueResponse, err *AuthleteError) {
	res = &dto.BackchannelAuthenticationIssueResponse{}
	err = self.callServicePostApi(`/api/backchannel/authentication/issue`, request, res)
	return
}

func (self *impl) BackchannelAuthenticationFail(
	request *dto.BackchannelAuthenticationFailRequest) (res *dto.BackchannelAuthenticationFailResponse, err *AuthleteError) {
	res = &dto.BackchannelAuthenticationFailResponse{}
	err = self.callServicePostApi(`/api/backchannel/authentication/fail`, request, res)
	return
}

func (self *impl) BackchannelAuthenticationComplete(
	request *dto.BackchannelAuthenticationCompleteRequest) (res *dto.BackchannelAuthenticationCompleteResponse, err *AuthleteError) {
	res = &dto.BackchannelAuthenticationCompleteResponse{}
	err = self.callServicePostApi(`/api/backchannel/authentication/complete`, request, res)
	return
}

func (self *impl) DeviceAuthorization(
	request *dto.DeviceAuthorizationRequest) (res *dto.DeviceAuthorizationResponse, err *AuthleteError) {
	res = &dto.DeviceAuthorizationResponse{}
	err = self.callServicePostApi(`/api/device/authorization`, request, res)
	return
}

func (self *impl) DeviceComplete(
	request *dto.DeviceCompleteRequest) (res *dto.DeviceCompleteResponse, err *AuthleteError) {
	res = &dto.DeviceCompleteResponse{}
	err = self.callServicePostApi(`/api/device/complete`, request, res)
	return
}

func (self *impl) DeviceVerification(
	request *dto.DeviceVerificationRequest) (res *dto.DeviceVerificationResponse, err *AuthleteError) {
	res = &dto.DeviceVerificationResponse{}
	err = self.callServicePostApi(`/api/device/verification`, request, res)
	return
}

func (self *impl) PushAuthorizationRequest(
	request *dto.PushedAuthReqRequest) (res *dto.PushedAuthReqResponse, err *AuthleteError) {
	res = &dto.PushedAuthReqResponse{}
	err = self.callServicePostApi(`/api/pushed_auth_req`, request, res)
	return
}

func (self *impl) HskCreate(
	request *dto.HskCreateRequest) (res *dto.HskResponse, err *AuthleteError) {
	res = &dto.HskResponse{}
	err = self.callServicePostApi(`/api/hsk/create`, request, res)
	return
}

func (self *impl) HskDelete(
	handle interface{}) (res *dto.HskResponse, err *AuthleteError) {
	path := `/api/hsk/delete/` + toString(handle)
	res = &dto.HskResponse{}
	err = self.callServiceGetApi(path, nil, res)
	return
}

func (self *impl) HskGet(
	handle interface{}) (res *dto.HskResponse, err *AuthleteError) {
	path := `/api/hsk/get/` + toString(handle)
	res = &dto.HskResponse{}
	err = self.callServiceGetApi(path, nil, res)
	return
}

func (self *impl) HskGetList() (
	res *dto.HskListResponse, err *AuthleteError) {
	res = &dto.HskListResponse{}
	err = self.callServiceGetApi(`/api/hsk/get/list`, nil, res)
	return
}

func New(configuration conf.AuthleteConfiguration) AuthleteApi {
	im := impl{}
	im.init(configuration)

	return &im
}
