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
	"context"
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
	u, err := url.Parse(configuration.GetBaseUrl())
	if err != nil {
		log.Panicf(`Invalid base URL '%s'`, configuration.GetBaseUrl())
	}
	self.baseUrl = strings.TrimRight(u.String(), `/`)
	self.serviceOwnerApiKey = configuration.GetServiceOwnerApiKey()
	self.serviceOwnerApiSecret = configuration.GetServiceOwnerApiSecret()
	self.serviceApiKey = configuration.GetServiceApiKey()
	self.serviceApiSecret = configuration.GetServiceApiSecret()
	self.settings = &Settings{}
}

// this method copy the raw response body from api server to
// better report error to clients using your api
func drainBody(body *io.ReadCloser) string {
	buf := new(strings.Builder)
	io.Copy(buf, *body)
	return buf.String()
}

func (self *impl) callApi(
	method string, apiKey string, apiSecret string, path string,
	queryParams map[string]string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	return self.callApiContext(
		context.Background(),
		method, apiKey, apiSecret, path,
		queryParams, requestBody,
		responseContainer)
}

func (self *impl) callApiContext(
	ctx context.Context,
	method string, apiKey string, apiSecret string, path string,
	queryParams map[string]string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	// Prepare a request to the Authlete API.
	req := self.buildRequest(method, apiKey, apiSecret, path, queryParams, requestBody).WithContext(ctx)

	// HTTP Client
	client := self.prepareClient()

	// Call the Authlete API.
	res, err := client.Do(req)
	if err != nil {
		msg := fmt.Sprintf(`API call to '%s' failed, with error %s`, path, err.Error())
		return &AuthleteError{err, msg, req, res}
	}

	// HTTP Status Code of the API response.
	sc := res.StatusCode
	if sc < 200 || 300 <= sc {
		msg := fmt.Sprintf(`'%s' API returned %d with body \n %s`, path, sc, drainBody(&res.Body))
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

	if len(apiKey) > 0 || len(apiSecret) > 0 {
		req.SetBasicAuth(apiKey, apiSecret)
	}

	return &req
}

func (self *impl) buildRequestHeader() http.Header {
	headers := http.Header{
		`Accept`:       {`application/json`},
		`Content-Type`: {`application/json;charset=UTF-8`},
	}
	if len(self.settings.UserAgent) > 0 {
		headers.Set(`User-Agent`, self.settings.UserAgent)
	}
	return headers
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
		Timeout:   self.settings.Timeout,
		Transport: self.settings.Transport,
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

func (self *impl) callGetApiContext(
	ctx context.Context,
	apiKey string, apiSecret string, path string, queryParams map[string]string,
	responseContainer interface{}) *AuthleteError {
	return self.callApiContext(
		ctx,
		http.MethodGet, apiKey, apiSecret,
		path, queryParams, nil, responseContainer)
}

func (self *impl) callGetApiWithoutCredentialsContext(
	ctx context.Context,
	path string, queryParams map[string]string, responseContainer interface{}) *AuthleteError {
	return self.callApiContext(
		ctx,
		http.MethodGet, "", "",
		path, queryParams, nil, responseContainer)
}

func (self *impl) callPostApiContext(
	ctx context.Context,
	apiKey string, apiSecret string, path string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	return self.callApiContext(
		ctx,
		http.MethodPost, apiKey, apiSecret,
		path, nil, requestBody, responseContainer)
}

func (self *impl) callDeleteApiContext(
	ctx context.Context,
	apiKey string, apiSecret string, path string) *AuthleteError {
	return self.callApiContext(
		ctx,
		http.MethodDelete, apiKey, apiSecret, path, nil, nil, nil)
}

func (self *impl) callServiceGetApiContext(
	ctx context.Context,
	path string, queryParams map[string]string,
	responseContainer interface{}) *AuthleteError {
	return self.callGetApiContext(
		ctx,
		self.serviceApiKey, self.serviceApiSecret,
		path, queryParams, responseContainer)
}

func (self *impl) callServicePostApiContext(
	ctx context.Context,
	path string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	return self.callPostApiContext(
		ctx,
		self.serviceApiKey, self.serviceApiSecret,
		path, requestBody, responseContainer)
}

func (self *impl) callServiceDeleteApicontext(ctx context.Context, path string) *AuthleteError {
	return self.callDeleteApiContext(
		ctx,
		self.serviceApiKey, self.serviceApiSecret, path)
}

func (self *impl) callServiceOwnerGetApiContext(
	ctx context.Context,
	path string, queryParams map[string]string,
	responseContainer interface{}) *AuthleteError {
	return self.callGetApiContext(
		ctx,
		self.serviceOwnerApiKey, self.serviceOwnerApiSecret,
		path, queryParams, responseContainer)
}

func (self *impl) callServiceOwnerPostApiContext(
	ctx context.Context,
	path string, requestBody interface{},
	responseContainer interface{}) *AuthleteError {
	return self.callPostApiContext(
		ctx,
		self.serviceOwnerApiKey, self.serviceOwnerApiSecret,
		path, requestBody, responseContainer)
}

func (self *impl) callServiceOwnerDeleteApiContext(ctx context.Context, path string) *AuthleteError {
	return self.callDeleteApiContext(
		ctx,
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
	return self.AuthorizationContext(context.Background(), request)
}

func (self *impl) AuthorizationContext(
	ctx context.Context,
	request *dto.AuthorizationRequest) (res *dto.AuthorizationResponse, err *AuthleteError) {
	res = &dto.AuthorizationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/authorization`, request, res)
	return
}

func (self *impl) AuthorizationFail(
	request *dto.AuthorizationFailRequest) (res *dto.AuthorizationFailResponse, err *AuthleteError) {
	return self.AuthorizationFailContext(context.Background(), request)
}

func (self *impl) AuthorizationFailContext(
	ctx context.Context,
	request *dto.AuthorizationFailRequest) (res *dto.AuthorizationFailResponse, err *AuthleteError) {
	res = &dto.AuthorizationFailResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/authorization/fail`, request, res)
	return
}

func (self *impl) AuthorizationIssue(
	request *dto.AuthorizationIssueRequest) (res *dto.AuthorizationIssueResponse, err *AuthleteError) {
	return self.AuthorizationIssueContext(context.Background(), request)
}

func (self *impl) AuthorizationIssueContext(
	ctx context.Context,
	request *dto.AuthorizationIssueRequest) (res *dto.AuthorizationIssueResponse, err *AuthleteError) {
	res = &dto.AuthorizationIssueResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/authorization/issue`, request, res)
	return
}

func (self *impl) Token(
	request *dto.TokenRequest) (res *dto.TokenResponse, err *AuthleteError) {
	return self.TokenContext(context.Background(), request)
}

func (self *impl) TokenContext(
	ctx context.Context,
	request *dto.TokenRequest) (res *dto.TokenResponse, err *AuthleteError) {
	res = &dto.TokenResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/token`, request, res)
	return
}

func (self *impl) TokenCreate(
	request *dto.TokenCreateRequest) (res *dto.TokenCreateResponse, err *AuthleteError) {
	return self.TokenCreateContext(context.Background(), request)
}

func (self *impl) TokenCreateContext(
	ctx context.Context,
	request *dto.TokenCreateRequest) (res *dto.TokenCreateResponse, err *AuthleteError) {
	res = &dto.TokenCreateResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/token/create`, request, res)
	return
}

func (self *impl) TokenDelete(
	token string) (err *AuthleteError) {
	return self.TokenDeleteContext(context.Background(), token)
}

func (self *impl) TokenDeleteContext(
	ctx context.Context,
	token string) (err *AuthleteError) {
	path := `/api/auth/token/delete/` + token
	err = self.callServiceDeleteApicontext(ctx, path)
	return
}

func (self *impl) TokenFail(
	request *dto.TokenFailRequest) (res *dto.TokenFailResponse, err *AuthleteError) {
	return self.TokenFailContext(context.Background(), request)
}

func (self *impl) TokenFailContext(
	ctx context.Context,
	request *dto.TokenFailRequest) (res *dto.TokenFailResponse, err *AuthleteError) {
	res = &dto.TokenFailResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/token/fail`, request, res)
	return
}

func (self *impl) TokenIssue(
	request *dto.TokenIssueRequest) (res *dto.TokenIssueResponse, err *AuthleteError) {
	return self.TokenIssueContext(context.Background(), request)
}

func (self *impl) TokenIssueContext(
	ctx context.Context,
	request *dto.TokenIssueRequest) (res *dto.TokenIssueResponse, err *AuthleteError) {
	res = &dto.TokenIssueResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/token/issue`, request, res)
	return
}

func (self *impl) TokenUpdate(
	request *dto.TokenUpdateRequest) (res *dto.TokenUpdateResponse, err *AuthleteError) {
	return self.TokenUpdateContext(context.Background(), request)
}

func (self *impl) TokenUpdateContext(
	ctx context.Context,
	request *dto.TokenUpdateRequest) (res *dto.TokenUpdateResponse, err *AuthleteError) {
	res = &dto.TokenUpdateResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/token/update`, request, res)
	return
}

func (self *impl) GetTokenList(
	clientIdentifier string, subject string, start uint32, end uint32) (res *dto.TokenListResponse, err *AuthleteError) {
	return self.GetTokenListContext(context.Background(), clientIdentifier, subject, start, end)
}
func (self *impl) GetTokenListContext(
	ctx context.Context,
	clientIdentifier string, subject string, start uint32, end uint32) (res *dto.TokenListResponse, err *AuthleteError) {
	params := buildMap(
		`clientIdentifier`, clientIdentifier,
		`subject`, subject, `start`, start, `end`, end)

	res = &dto.TokenListResponse{}
	err = self.callServiceGetApiContext(ctx, `/api/auth/token/get/list`, params, res)
	return
}

func (self *impl) Revocation(
	request *dto.RevocationRequest) (res *dto.RevocationResponse, err *AuthleteError) {
	return self.RevocationContext(context.Background(), request)
}

func (self *impl) RevocationContext(
	ctx context.Context,
	request *dto.RevocationRequest) (res *dto.RevocationResponse, err *AuthleteError) {
	res = &dto.RevocationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/revocation`, request, res)
	return
}

func (self *impl) UserInfo(
	request *dto.UserInfoRequest) (res *dto.UserInfoResponse, err *AuthleteError) {
	return self.UserInfoContext(context.Background(), request)
}

func (self *impl) UserInfoContext(
	ctx context.Context,
	request *dto.UserInfoRequest) (res *dto.UserInfoResponse, err *AuthleteError) {
	res = &dto.UserInfoResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/userinfo`, request, res)
	return
}

func (self *impl) UserInfoIssue(
	request *dto.UserInfoIssueRequest) (res *dto.UserInfoIssueResponse, err *AuthleteError) {
	return self.UserInfoIssueContext(context.Background(), request)
}

func (self *impl) UserInfoIssueContext(
	ctx context.Context,
	request *dto.UserInfoIssueRequest) (res *dto.UserInfoIssueResponse, err *AuthleteError) {
	res = &dto.UserInfoIssueResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/userinfo/issue`, request, res)
	return
}

func (self *impl) Introspection(
	request *dto.IntrospectionRequest) (res *dto.IntrospectionResponse, err *AuthleteError) {
	return self.IntrospectionContext(context.Background(), request)
}

func (self *impl) IntrospectionContext(
	ctx context.Context,
	request *dto.IntrospectionRequest) (res *dto.IntrospectionResponse, err *AuthleteError) {
	res = &dto.IntrospectionResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/introspection`, request, res)
	return
}

func (self *impl) StandardIntrospection(
	request *dto.StandardIntrospectionRequest) (res *dto.StandardIntrospectionResponse, err *AuthleteError) {
	return self.StandardIntrospectionContext(context.Background(), request)
}

func (self *impl) StandardIntrospectionContext(
	ctx context.Context,
	request *dto.StandardIntrospectionRequest) (res *dto.StandardIntrospectionResponse, err *AuthleteError) {
	res = &dto.StandardIntrospectionResponse{}
	err = self.callServicePostApiContext(ctx, `/api/auth/introspection/standard`, request, res)
	return
}

func (self *impl) CreateService(
	request *dto.Service) (res *dto.Service, err *AuthleteError) {
	return self.CreateServiceContext(context.Background(), request)
}

func (self *impl) CreateServiceContext(
	ctx context.Context,
	request *dto.Service) (res *dto.Service, err *AuthleteError) {
	res = &dto.Service{}
	err = self.callServiceOwnerPostApiContext(ctx, `/api/service/create`, request, res)
	return
}

func (self *impl) DeleteService(
	apiKey interface{}) (err *AuthleteError) {
	return self.DeleteServiceContext(context.Background(), apiKey)
}

func (self *impl) DeleteServiceContext(
	ctx context.Context,
	apiKey interface{}) (err *AuthleteError) {
	path := `/api/service/delete/` + toString(apiKey)
	err = self.callServiceOwnerDeleteApiContext(ctx, path)
	return
}

func (self *impl) GetService(
	apiKey interface{}) (res *dto.Service, err *AuthleteError) {
	return self.GetServiceContext(context.Background(), apiKey)
}

func (self *impl) GetServiceContext(
	ctx context.Context,
	apiKey interface{}) (res *dto.Service, err *AuthleteError) {
	path := `/api/service/get/` + toString(apiKey)
	res = &dto.Service{}
	err = self.callServiceOwnerGetApiContext(ctx, path, nil, res)
	return
}

func (self *impl) GetServiceList(
	start uint32, end uint32) (res *dto.ServiceListResponse, err *AuthleteError) {
	return self.GetServiceListContext(context.Background(), start, end)
}

func (self *impl) GetServiceListContext(
	ctx context.Context,
	start uint32, end uint32) (res *dto.ServiceListResponse, err *AuthleteError) {
	params := buildMap(`start`, start, `end`, end)

	res = &dto.ServiceListResponse{}
	err = self.callServiceOwnerGetApiContext(ctx, `/api/service/get/list`, params, res)
	return
}

func (self *impl) UpdateService(
	service *dto.Service) (res *dto.Service, err *AuthleteError) {
	return self.UpdateServiceContext(context.Background(), service)
}

func (self *impl) UpdateServiceContext(
	ctx context.Context,
	service *dto.Service) (res *dto.Service, err *AuthleteError) {
	path := `/api/service/update/` + toString(service.ApiKey)
	res = &dto.Service{}
	err = self.callServiceOwnerPostApiContext(ctx, path, service, res)
	return
}

func (self *impl) GetServiceJwks(
	pretty bool, includePrivateKeys bool) (res string, err *AuthleteError) {
	return self.GetServiceJwksContext(context.Background(), pretty, includePrivateKeys)
}

func (self *impl) GetServiceJwksContext(
	ctx context.Context,
	pretty bool, includePrivateKeys bool) (res string, err *AuthleteError) {
	params := buildMap(`pretty`, pretty, `includePrivateKeys`, includePrivateKeys)

	obj := strings.Builder{}
	err = self.callServiceGetApiContext(ctx, `/api/service/jwks/get`, params, &obj)

	if err == nil {
		res = obj.String()
	}

	return
}

func (self *impl) GetServiceConfiguration(
	pretty bool) (res string, err *AuthleteError) {
	return self.GetServiceConfigurationContext(context.Background(), pretty)
}

func (self *impl) GetServiceConfigurationContext(
	ctx context.Context,
	pretty bool) (res string, err *AuthleteError) {
	params := buildMap(`pretty`, pretty)

	obj := strings.Builder{}
	err = self.callServiceGetApiContext(ctx, `/api/service/configuration`, params, &obj)

	if err == nil {
		res = obj.String()
	}

	return
}

func (self *impl) CreateClient(
	client *dto.Client) (res *dto.Client, err *AuthleteError) {
	return self.CreateClientContext(context.Background(), client)
}

func (self *impl) CreateClientContext(
	ctx context.Context,
	client *dto.Client) (res *dto.Client, err *AuthleteError) {
	res = &dto.Client{}
	err = self.callServicePostApiContext(ctx, `/api/client/create`, client, res)
	return
}

func (self *impl) DynamicClientRegister(
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	return self.DynamicClientRegisterContext(context.Background(), request)
}

func (self *impl) DynamicClientRegisterContext(
	ctx context.Context,
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	res = &dto.ClientRegistrationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/client/registration`, request, res)
	return
}

func (self *impl) DynamicClientGet(
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	return self.DynamicClientGetContext(context.Background(), request)
}

func (self *impl) DynamicClientGetContext(
	ctx context.Context,
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	res = &dto.ClientRegistrationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/client/registration/get`, request, res)
	return
}

func (self *impl) DynamicClientUpdate(
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	return self.DynamicClientUpdateContext(context.Background(), request)
}

func (self *impl) DynamicClientUpdateContext(
	ctx context.Context,
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	res = &dto.ClientRegistrationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/client/registration/update`, request, res)
	return
}

func (self *impl) DynamicClientDelete(
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	return self.DynamicClientDeleteContext(context.Background(), request)
}

func (self *impl) DynamicClientDeleteContext(
	ctx context.Context,
	request *dto.ClientRegistrationRequest) (res *dto.ClientRegistrationResponse, err *AuthleteError) {
	res = &dto.ClientRegistrationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/client/registration/delete`, request, res)
	return
}

func (self *impl) DeleteClient(
	clientIdentifier interface{}) (err *AuthleteError) {
	return self.DeleteClientContext(context.Background(), clientIdentifier)
}

func (self *impl) DeleteClientContext(
	ctx context.Context,
	clientIdentifier interface{}) (err *AuthleteError) {
	path := `/api/client/delete/` + toString(clientIdentifier)
	err = self.callServiceDeleteApicontext(ctx, path)
	return
}

func (self *impl) GetClient(
	clientIdentifier interface{}) (res *dto.Client, err *AuthleteError) {
	return self.GetClientContext(context.Background(), clientIdentifier)
}

func (self *impl) GetClientContext(
	ctx context.Context,
	clientIdentifier interface{}) (res *dto.Client, err *AuthleteError) {
	path := `/api/client/get/` + toString(clientIdentifier)
	res = &dto.Client{}
	err = self.callServiceGetApiContext(ctx, path, nil, res)
	return
}

func (self *impl) GetClientList(
	developer string, start uint32, end uint32) (res *dto.ClientListResponse, err *AuthleteError) {
	return self.GetClientListContext(context.Background(), developer, start, end)
}

func (self *impl) GetClientListContext(
	ctx context.Context,
	developer string, start uint32, end uint32) (res *dto.ClientListResponse, err *AuthleteError) {
	params := buildMap(`developer`, developer, `start`, start, `end`, end)

	res = &dto.ClientListResponse{}
	err = self.callServiceGetApiContext(ctx, `/api/client/get/list`, params, res)
	return
}

func (self *impl) UpdateClient(
	client *dto.Client) (res *dto.Client, err *AuthleteError) {
	return self.UpdateClientContext(context.Background(), client)
}

func (self *impl) UpdateClientContext(
	ctx context.Context,
	client *dto.Client) (res *dto.Client, err *AuthleteError) {
	path := `/api/client/update/` + toString(client.ClientId)
	res = &dto.Client{}
	err = self.callServicePostApiContext(ctx, path, client, res)
	return
}

type requestableScopes struct {
	RequestableScopes []string `json:"requestableScopes,omitempty"`
}

func (self *impl) GetRequestableScopes(
	clientIdentifier interface{}) (res []string, err *AuthleteError) {
	return self.GetRequestableScopesContext(context.Background(), clientIdentifier)
}

func (self *impl) GetRequestableScopesContext(
	ctx context.Context,
	clientIdentifier interface{}) (res []string, err *AuthleteError) {
	path := `/api/client/extension/requestable_scopes/get/` + toString(clientIdentifier)

	obj := requestableScopes{}
	err = self.callServiceGetApiContext(ctx, path, nil, &obj)

	if err == nil {
		res = obj.RequestableScopes
	}

	return
}

func (self *impl) SetRequestableScopes(
	clientIdentifier interface{}, scopes []string) (res []string, err *AuthleteError) {
	return self.SetRequestableScopesContext(context.Background(), clientIdentifier, scopes)
}

func (self *impl) SetRequestableScopesContext(
	ctx context.Context,
	clientIdentifier interface{}, scopes []string) (res []string, err *AuthleteError) {
	input := requestableScopes{}
	input.RequestableScopes = scopes

	path := `/api/client/extension/requestable_scopes/update/` + toString(clientIdentifier)

	output := requestableScopes{}
	err = self.callServicePostApiContext(ctx, path, &input, &output)

	if err == nil {
		res = output.RequestableScopes
	}

	return
}

func (self *impl) DeleteRequestableScopes(
	clientIdentifier interface{}) (err *AuthleteError) {
	return self.DeleteRequestableScopesContext(context.Background(), clientIdentifier)
}

func (self *impl) DeleteRequestableScopesContext(
	ctx context.Context,
	clientIdentifier interface{}) (err *AuthleteError) {
	path := `/api/client/extension/requestable_scopes/delete/` + toString(clientIdentifier)
	err = self.callServiceDeleteApicontext(ctx, path)
	return
}

type grantedScopesRequest struct {
	Subject string `json:"subject,omitempty"`
}

func (self *impl) GetGrantedScopes(
	clientIdentifier interface{}, subject string) (res *dto.GrantedScopesGetResponse, err *AuthleteError) {
	return self.GetGrantedScopesContext(context.Background(), clientIdentifier, subject)
}

func (self *impl) GetGrantedScopesContext(
	ctx context.Context,
	clientIdentifier interface{}, subject string) (res *dto.GrantedScopesGetResponse, err *AuthleteError) {
	req := grantedScopesRequest{}
	req.Subject = subject

	path := `/api/client/granted_scopes/get/` + toString(clientIdentifier)

	res = &dto.GrantedScopesGetResponse{}
	err = self.callServicePostApiContext(ctx, path, &req, res)
	return
}

func (self *impl) DeleteGrantedScopes(
	clientIdentifier interface{}, subject string) (err *AuthleteError) {
	return self.DeleteGrantedScopesContext(context.Background(), clientIdentifier, subject)
}

func (self *impl) DeleteGrantedScopesContext(
	ctx context.Context,
	clientIdentifier interface{}, subject string) (err *AuthleteError) {
	path := `/api/client/granted_scopes/delete/` + toString(clientIdentifier)
	err = self.callServiceDeleteApicontext(ctx, path)
	return
}

func (self *impl) DeleteClientAuthorization(
	clientIdentifier interface{}, subject string) (err *AuthleteError) {
	return self.DeleteClientAuthorizationContext(context.Background(), clientIdentifier, subject)
}

func (self *impl) DeleteClientAuthorizationContext(
	ctx context.Context,
	clientIdentifier interface{}, subject string) (err *AuthleteError) {
	req := dto.ClientAuthorizationDeleteRequest{}
	req.Subject = subject

	path := `/api/client/authorization/delete/` + toString(clientIdentifier)

	res := &dto.ApiResponse{}
	err = self.callServicePostApiContext(ctx, path, &req, res)
	return
}

func (self *impl) GetClientAuthorizationList(
	request *dto.ClientAuthorizationGetListRequest) (res *dto.AuthorizedClientListResponse, err *AuthleteError) {
	return self.GetClientAuthorizationListContext(context.Background(), request)
}

func (self *impl) GetClientAuthorizationListContext(
	ctx context.Context,
	request *dto.ClientAuthorizationGetListRequest) (res *dto.AuthorizedClientListResponse, err *AuthleteError) {
	res = &dto.AuthorizedClientListResponse{}
	err = self.callServicePostApiContext(ctx, `/api/client/authorization/get/list`, request, res)
	return
}

func (self *impl) UpdateClientAuthorization(
	clientIdentifier interface{}, request *dto.ClientAuthorizationUpdateRequest) (err *AuthleteError) {
	return self.UpdateClientAuthorizationContext(context.Background(), clientIdentifier, request)
}

func (self *impl) UpdateClientAuthorizationContext(
	ctx context.Context,
	clientIdentifier interface{}, request *dto.ClientAuthorizationUpdateRequest) (err *AuthleteError) {
	path := `/api/client/authorization/update/` + toString(clientIdentifier)

	res := &dto.ApiResponse{}
	err = self.callServicePostApiContext(ctx, path, request, res)
	return
}

func (self *impl) RefreshClientSecret(
	clientIdentifier interface{}) (res *dto.ClientSecretRefreshResponse, err *AuthleteError) {
	return self.RefreshClientSecretContext(context.Background(), clientIdentifier)
}

func (self *impl) RefreshClientSecretContext(
	ctx context.Context,
	clientIdentifier interface{}) (res *dto.ClientSecretRefreshResponse, err *AuthleteError) {
	path := `/api/client/secret/refresh/` + toString(clientIdentifier)

	res = &dto.ClientSecretRefreshResponse{}
	err = self.callServiceGetApiContext(ctx, path, nil, res)
	return
}

func (self *impl) UpdateClientSecret(
	clientIdentifier interface{}, clientSecret string) (res *dto.ClientSecretUpdateResponse, err *AuthleteError) {
	return self.UpdateClientSecretContext(context.Background(), clientIdentifier, clientSecret)
}

func (self *impl) UpdateClientSecretContext(
	ctx context.Context,
	clientIdentifier interface{}, clientSecret string) (res *dto.ClientSecretUpdateResponse, err *AuthleteError) {
	req := dto.ClientSecretUpdateRequest{}
	req.ClientSecret = clientSecret

	path := `/api/client/secret/update/` + toString(clientIdentifier)

	res = &dto.ClientSecretUpdateResponse{}
	err = self.callServicePostApiContext(ctx, path, &req, res)
	return
}

func (self *impl) VerifyJose(
	request *dto.JoseVerifyRequest) (res *dto.JoseVerifyResponse, err *AuthleteError) {
	return self.VerifyJoseContext(context.Background(), request)
}

func (self *impl) VerifyJoseContext(
	ctx context.Context,
	request *dto.JoseVerifyRequest) (res *dto.JoseVerifyResponse, err *AuthleteError) {
	res = &dto.JoseVerifyResponse{}
	err = self.callServicePostApiContext(ctx, `/api/jose/verify`, request, res)
	return
}

func (self *impl) BackchannelAuthentication(
	request *dto.BackchannelAuthenticationRequest) (res *dto.BackchannelAuthenticationResponse, err *AuthleteError) {
	return self.BackchannelAuthenticationContext(context.Background(), request)
}

func (self *impl) BackchannelAuthenticationContext(
	ctx context.Context,
	request *dto.BackchannelAuthenticationRequest) (res *dto.BackchannelAuthenticationResponse, err *AuthleteError) {
	res = &dto.BackchannelAuthenticationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/backchannel/authentication`, request, res)
	return
}

func (self *impl) BackchannelAuthenticationIssue(
	request *dto.BackchannelAuthenticationIssueRequest) (res *dto.BackchannelAuthenticationIssueResponse, err *AuthleteError) {
	return self.BackchannelAuthenticationIssueContext(context.Background(), request)
}

func (self *impl) BackchannelAuthenticationIssueContext(
	ctx context.Context,
	request *dto.BackchannelAuthenticationIssueRequest) (res *dto.BackchannelAuthenticationIssueResponse, err *AuthleteError) {
	res = &dto.BackchannelAuthenticationIssueResponse{}
	err = self.callServicePostApiContext(ctx, `/api/backchannel/authentication/issue`, request, res)
	return
}

func (self *impl) BackchannelAuthenticationFail(
	request *dto.BackchannelAuthenticationFailRequest) (res *dto.BackchannelAuthenticationFailResponse, err *AuthleteError) {
	return self.BackchannelAuthenticationFailContext(context.Background(), request)
}

func (self *impl) BackchannelAuthenticationFailContext(
	ctx context.Context,
	request *dto.BackchannelAuthenticationFailRequest) (res *dto.BackchannelAuthenticationFailResponse, err *AuthleteError) {
	res = &dto.BackchannelAuthenticationFailResponse{}
	err = self.callServicePostApiContext(ctx, `/api/backchannel/authentication/fail`, request, res)
	return
}

func (self *impl) BackchannelAuthenticationComplete(
	request *dto.BackchannelAuthenticationCompleteRequest) (res *dto.BackchannelAuthenticationCompleteResponse, err *AuthleteError) {
	return self.BackchannelAuthenticationCompleteContext(context.Background(), request)
}

func (self *impl) BackchannelAuthenticationCompleteContext(
	ctx context.Context,
	request *dto.BackchannelAuthenticationCompleteRequest) (res *dto.BackchannelAuthenticationCompleteResponse, err *AuthleteError) {
	res = &dto.BackchannelAuthenticationCompleteResponse{}
	err = self.callServicePostApiContext(ctx, `/api/backchannel/authentication/complete`, request, res)
	return
}

func (self *impl) DeviceAuthorization(
	request *dto.DeviceAuthorizationRequest) (res *dto.DeviceAuthorizationResponse, err *AuthleteError) {
	return self.DeviceAuthorizationContext(context.Background(), request)
}

func (self *impl) DeviceAuthorizationContext(
	ctx context.Context,
	request *dto.DeviceAuthorizationRequest) (res *dto.DeviceAuthorizationResponse, err *AuthleteError) {
	res = &dto.DeviceAuthorizationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/device/authorization`, request, res)
	return
}

func (self *impl) DeviceComplete(
	request *dto.DeviceCompleteRequest) (res *dto.DeviceCompleteResponse, err *AuthleteError) {
	return self.DeviceCompleteContext(context.Background(), request)
}

func (self *impl) DeviceCompleteContext(
	ctx context.Context,
	request *dto.DeviceCompleteRequest) (res *dto.DeviceCompleteResponse, err *AuthleteError) {
	res = &dto.DeviceCompleteResponse{}
	err = self.callServicePostApiContext(ctx, `/api/device/complete`, request, res)
	return
}

func (self *impl) DeviceVerification(
	request *dto.DeviceVerificationRequest) (res *dto.DeviceVerificationResponse, err *AuthleteError) {
	return self.DeviceVerificationContext(context.Background(), request)
}

func (self *impl) DeviceVerificationContext(
	ctx context.Context,
	request *dto.DeviceVerificationRequest) (res *dto.DeviceVerificationResponse, err *AuthleteError) {
	res = &dto.DeviceVerificationResponse{}
	err = self.callServicePostApiContext(ctx, `/api/device/verification`, request, res)
	return
}

func (self *impl) PushAuthorizationRequest(
	request *dto.PushedAuthReqRequest) (res *dto.PushedAuthReqResponse, err *AuthleteError) {
	return self.PushAuthorizationRequestContext(context.Background(), request)
}

func (self *impl) PushAuthorizationRequestContext(
	ctx context.Context,
	request *dto.PushedAuthReqRequest) (res *dto.PushedAuthReqResponse, err *AuthleteError) {
	res = &dto.PushedAuthReqResponse{}
	err = self.callServicePostApiContext(ctx, `/api/pushed_auth_req`, request, res)
	return
}

func (self *impl) HskCreate(
	request *dto.HskCreateRequest) (res *dto.HskResponse, err *AuthleteError) {
	return self.HskCreateContext(context.Background(), request)
}

func (self *impl) HskCreateContext(
	ctx context.Context,
	request *dto.HskCreateRequest) (res *dto.HskResponse, err *AuthleteError) {
	res = &dto.HskResponse{}
	err = self.callServicePostApiContext(ctx, `/api/hsk/create`, request, res)
	return
}

func (self *impl) HskDelete(
	handle interface{}) (res *dto.HskResponse, err *AuthleteError) {
	return self.HskDeleteContext(context.Background(), handle)
}

func (self *impl) HskDeleteContext(
	ctx context.Context,
	handle interface{}) (res *dto.HskResponse, err *AuthleteError) {
	path := `/api/hsk/delete/` + toString(handle)
	res = &dto.HskResponse{}
	err = self.callServiceGetApiContext(ctx, path, nil, res)
	return
}

func (self *impl) HskGet(
	handle interface{}) (res *dto.HskResponse, err *AuthleteError) {
	return self.HskGetContext(context.Background(), handle)
}

func (self *impl) HskGetContext(
	ctx context.Context,
	handle interface{}) (res *dto.HskResponse, err *AuthleteError) {
	path := `/api/hsk/get/` + toString(handle)
	res = &dto.HskResponse{}
	err = self.callServiceGetApiContext(ctx, path, nil, res)
	return
}

func (self *impl) HskGetList() (
	res *dto.HskListResponse, err *AuthleteError) {
	return self.HskGetListContext(context.Background())
}

func (self *impl) HskGetListContext(ctx context.Context) (
	res *dto.HskListResponse, err *AuthleteError) {
	res = &dto.HskListResponse{}
	err = self.callServiceGetApiContext(ctx, `/api/hsk/get/list`, nil, res)
	return
}

func (self *impl) Echo(parameters *map[string]string) (res *map[string]string, err *AuthleteError) {
	return self.EchoContext(context.Background(), parameters)
}

func (self *impl) EchoContext(ctx context.Context, parameters *map[string]string) (res *map[string]string, err *AuthleteError) {
	res = &map[string]string{}
	err = self.callGetApiWithoutCredentialsContext(ctx, `/api/misc/echo`, *parameters, res)
	return
}

func New(configuration conf.AuthleteConfiguration) AuthleteApi {
	im := impl{}
	im.init(configuration)

	return &im
}
