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

import "github.com/authlete/authlete-go/dto"

type AuthleteApi interface {
	Settings() *Settings
	Authorization(request *dto.AuthorizationRequest) (*dto.AuthorizationResponse, *AuthleteError)
	AuthorizationFail(request *dto.AuthorizationFailRequest) (*dto.AuthorizationFailResponse, *AuthleteError)
	AuthorizationIssue(request *dto.AuthorizationIssueRequest) (*dto.AuthorizationIssueResponse, *AuthleteError)
	Token(request *dto.TokenRequest) (*dto.TokenResponse, *AuthleteError)
	TokenCreate(request *dto.TokenCreateRequest) (*dto.TokenCreateResponse, *AuthleteError)
	TokenDelete(token string) *AuthleteError
	TokenFail(request *dto.TokenFailRequest) (*dto.TokenFailResponse, *AuthleteError)
	TokenIssue(request *dto.TokenIssueRequest) (*dto.TokenIssueResponse, *AuthleteError)
	TokenUpdate(request *dto.TokenUpdateRequest) (*dto.TokenUpdateResponse, *AuthleteError)
	GetTokenList(clientIdentifier string, subject string, start uint32, end uint32) (*dto.TokenListResponse, *AuthleteError)
	Revocation(request *dto.RevocationRequest) (*dto.RevocationResponse, *AuthleteError)
	UserInfo(request *dto.UserInfoRequest) (*dto.UserInfoResponse, *AuthleteError)
	UserInfoIssue(request *dto.UserInfoIssueRequest) (*dto.UserInfoIssueResponse, *AuthleteError)
	Introspection(request *dto.IntrospectionRequest) (*dto.IntrospectionResponse, *AuthleteError)
	StandardIntrospection(request *dto.StandardIntrospectionRequest) (*dto.StandardIntrospectionResponse, *AuthleteError)
	CreateService(service *dto.Service) (*dto.Service, *AuthleteError)
	DeleteService(apiKey interface{}) *AuthleteError
	GetService(apiKey interface{}) (*dto.Service, *AuthleteError)
	GetServiceList(start uint32, end uint32) (*dto.ServiceListResponse, *AuthleteError)
	UpdateService(service *dto.Service) (*dto.Service, *AuthleteError)
	GetServiceJwks(pretty bool, includePrivateKeys bool) (string, *AuthleteError)
	GetServiceConfiguration(pretty bool) (string, *AuthleteError)
	CreateClient(client *dto.Client) (*dto.Client, *AuthleteError)
	DynamicClientRegister(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)
	DynamicClientGet(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)
	DynamicClientUpdate(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)
	DynamicClientDelete(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)
	DeleteClient(clientIdentifier interface{}) *AuthleteError
	GetClient(clientIdentifier interface{}) (*dto.Client, *AuthleteError)
	GetClientList(developer string, start uint32, end uint32) (*dto.ClientListResponse, *AuthleteError)
	UpdateClient(client *dto.Client) (*dto.Client, *AuthleteError)
	GetRequestableScopes(clientIdentifier interface{}) ([]string, *AuthleteError)
	SetRequestableScopes(clientIdentifier interface{}, scopes []string) ([]string, *AuthleteError)
	DeleteRequestableScopes(clientIdentifier interface{}) *AuthleteError
	GetGrantedScopes(clientIdentifier interface{}, subject string) (*dto.GrantedScopesGetResponse, *AuthleteError)
	DeleteGrantedScopes(clientIdentifier interface{}, subject string) *AuthleteError
	DeleteClientAuthorization(clientIdentifier interface{}, subject string) *AuthleteError
	GetClientAuthorizationList(request *dto.ClientAuthorizationGetListRequest) (*dto.AuthorizedClientListResponse, *AuthleteError)
	UpdateClientAuthorization(clientIdentifier interface{}, request *dto.ClientAuthorizationUpdateRequest) *AuthleteError
	RefreshClientSecret(clientIdentifier interface{}) (*dto.ClientSecretRefreshResponse, *AuthleteError)
	UpdateClientSecret(clientIdentifier interface{}, clientSecret string) (*dto.ClientSecretUpdateResponse, *AuthleteError)
	VerifyJose(request *dto.JoseVerifyRequest) (*dto.JoseVerifyResponse, *AuthleteError)
	BackchannelAuthentication(request *dto.BackchannelAuthenticationRequest) (*dto.BackchannelAuthenticationResponse, *AuthleteError)
	BackchannelAuthenticationIssue(request *dto.BackchannelAuthenticationIssueRequest) (*dto.BackchannelAuthenticationIssueResponse, *AuthleteError)
	BackchannelAuthenticationFail(request *dto.BackchannelAuthenticationFailRequest) (*dto.BackchannelAuthenticationFailResponse, *AuthleteError)
	BackchannelAuthenticationComplete(request *dto.BackchannelAuthenticationCompleteRequest) (*dto.BackchannelAuthenticationCompleteResponse, *AuthleteError)
	DeviceAuthorization(request *dto.DeviceAuthorizationRequest) (*dto.DeviceAuthorizationResponse, *AuthleteError)
	DeviceComplete(request *dto.DeviceCompleteRequest) (*dto.DeviceCompleteResponse, *AuthleteError)
	DeviceVerification(request *dto.DeviceVerificationRequest) (*dto.DeviceVerificationResponse, *AuthleteError)
	PushAuthorizationRequest(request *dto.PushedAuthReqRequest) (*dto.PushedAuthReqResponse, *AuthleteError)
	HskCreate(request *dto.HskCreateRequest) (*dto.HskResponse, *AuthleteError)
	HskDelete(handle interface{}) (*dto.HskResponse, *AuthleteError)
	HskGet(handle interface{}) (*dto.HskResponse, *AuthleteError)
	HskGetList() (*dto.HskListResponse, *AuthleteError)
	Echo(parameters *map[string]string) (*map[string]string, *AuthleteError)
}
