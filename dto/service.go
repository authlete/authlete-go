//
// Copyright (C) 2019 Authlete, Inc.
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

package dto

import "github.com/authlete/authlete-go/types"

type Service struct {
	// The name of the service.
	ServiceName string `json:"serviceName"`

	// The API key of the service.
	ApiKey uint64 `json:"apiKey"`

	// The API secret of the service.
	ApiSecret string `json:"apiSecret"`

	// The issuer identifier.
	Issuer string `json:"issuer"`

	// The URL of the authorization endpoint.
	AuthorizationEndpoint string `json:"authorizationEndpoint"`

	// The URL of the token endpoint.
	TokenEndpoint string `json:"tokenEndpoint"`

	// The URL of the revocation endpoint.
	RevocationEndpoint string `json:"revocationEndpoint"`

	// Client authentication methods supported at the revocation endpoint.
	SupportedRevocationAuthMethods []types.ClientAuthMethod `json:"supportedRevocationAuthMethods"`

	// The URL of the userinfo endpoint.
	UserInfoEndpoint string `json:"userInfoEndpoint"`

	// The URI of the JWK Set document.
	JwksUri string `json:"jwksUri"`

	// The content of the JWK Set document.
	Jwks string `json:"jwks"`

	// The URL of the registration endpoint.
	RegistrationEndpoint string `json:"registrationEndpoint"`

	// The URL of the registration management endpoint.
	RegistrationManagementEndpoint string `json:"registrationManagementEndpoint"`

	// Supported scopes.
	Scopes []Scope `json:"scopes"`

	// Supported response types.
	SupportedResponseTypes []types.ResponseType `json:"supportedResponseTypes"`

	// Supported grant types.
	SupportedGrantTypes []types.GrantType `json:"supportedGrantTypes"`

	// String supported ACRs.
	SupportedAcrs []string `json:"supportedAcrs"`

	// Client authentication methods supported at the token endpoint.
	SupportedTokenAuthMethods []types.ClientAuthMethod `json:"supportedTokenAuthMethods"`

	// Supported displays.
	SupportedDisplays []types.Display `json:"supportedDisplays"`

	// Supported claim types.
	SupportedClaimTypes []types.ClaimType `json:"supportedClaimTypes"`

	// Supported claims.
	SupportedClaims []string `json:"supportedClaims"`

	// The URL of the service documentation.
	ServiceDocumentation string `json:"serviceDocumentation"`

	// Supported claim locales.
	SupportedClaimLocales []string `json:"supportedClaimLocales"`

	// Supported UI locales.
	SupportedUiLocales []string `json:"supportedUiLocales"`

	// The URL of the policy document.
	PolicyUri string `json:"policyUri"`

	// The URL of the terms of service.
	TosUri string `json:"tosUri"`

	// The URL of the authentication callback endpoint.
	AuthenticationCallbackEndpoint string `json:"authenticationCallbackEndpoint"`

	// The API key to access the authentication callback endpoint.
	AuthenticationCallbackApiKey string `json:"authenticationCallbackApiKey"`

	// The API secret to access the authentication callback endpoint.
	AuthenticationCallbackApiSecret string `json:"authenticationCallbackApiSecret"`

	// Supported SNSes.
	SupportedSnses []types.Sns `json:"supportedSnses"`

	// Credentials of supported SNSes.
	SnsCredentials []SnsCredentials `json:"snsCredentials"`

	// The URL of the developer authentication callback endpoint.
	DeveloperAuthenticationCallbackEndpoint string `json:"developerAuthenticationCallbackEndpoint"`

	// The API key to access the developer authentication callback endpoint.
	DeveloperAuthenticationCallbackApiKey string `json:"developerAuthenticationCallbackApiKey"`

	// The API secret to access the developer authentication callback endpoint.
	DeveloperAuthenticationCallbackApiSecret string `json:"developerAuthenticationCallbackApiSecret"`

	// The upper limit of the number of client applications per developer.
	ClientsPerDeveloper uint16 `json:"clientsPerDeveloper"`

	// The flag which indicates whether the direct authorization endpoint is enabled.
	DirectAuthorizationEndpointEnabled bool `json:"directAuthorizationEndpointEnabled"`

	// The flag which indicates whether the direct token endpoint is enabled.
	DirectTokenEndpointEnabled bool `json:"directTokenEndpointEnabled"`

	// The flag which indicates whether the direct revocation endpoint is enabled.
	DirectRevocationEndpointEnabled bool `json:"directRevocationEndpointEnabled"`

	// The flag which indicates whether the direct userinfo endpoint is enabled.
	DirectUserInfoEndpointEnabled bool `json:"directUserInfoEndpointEnabled"`

	// The flag which indicates whether the direct jwks endpoint is enabled.
	DirectJwksEndpointEnabled bool `json:"directJwksEndpointEnabled"`

	// The flag which indicates whether the direct introspection endpoint is enabled.
	DirectIntrospectionEndpointEnabled bool `json:"directIntrospectionEndpointEnabled"`

	// The flag which indicates whether the number of access tokens per subject is limited to at most 1.
	SingleAccessTokenPerSubject bool `json:"singleAccessTokenPerSubject"`

	// The flag which indicates whether PKCE is always required.
	PkceRequired bool `json:"pkceRequired"`

	// The flag which indicates whether S256 is always required for code_challenge_method.
	PkceS256Required bool `json:"pkceS256Required"`

	// The flag which indicates whether refresh tokens remain valid after use.
	RefreshTokenKept bool `json:"refreshTokenKept"`

	// The flag which indicates whether `error_description` is omitted.
	ErrorDescriptionOmitted bool `json:"errorDescriptionOmitted"`

	// The flag which indicates whether `error_uri` is omitted.
	ErrorUriOmitted bool `json:"errorUriOmitted"`

	// The flag which indicates whether the feature of Client ID Alias is enabled.
	ClientIdAliaseEnabled bool `json:"clientIdAliasEnabled"`

	// Supported service profiles.
	SupportedServiceProfiles []types.ServiceProfile `json:"supportedServiceProfiles"`

	// The flag which indicates whether certificate binding is supported.
	TlsClientCertificateBoundAccessTokens bool `json:"tlsClientCertificateBoundAccessTokens"`

	// The URL of the introspection endpoint.
	IntrospectionEndpoint string `json:"introspectionEndpoint"`

	// Client authentication methods supports at the introspection endpoint.
	SupportedIntrospectionAuthMethods []types.ClientAuthMethod `json:"supportedIntrospectionAuthMethods"`

	// The flag which indicates whether certification chain is validated for MTLS.
	MutualTlsValidatePkiCertChain bool `json:"mutualTlsValidatePkiCertChain"`

	// The list of trusted root certificates for MTLS.
	TrustedRootCertificates []string `json:"trustedRootCertificates"`

	// The flag which indicates whether Dynamic Client Registration is supported.
	DynamicRegistrationSupported bool `json:"dynamicRegistrationSupported"`

	// The description of this service.
	Description string `json:"description"`

	// The type of access token.
	AccessTokenType string `json:"accessTokenType"`

	// Signature algorithm of JWT-based access tokens.
	//
	// When this property is not nil, access tokens issued by this service are
	// JWS. Otherwise, access tokens are random strings as before.
	AccessTokenSignAlg types.JWSAlg `json:"accessTokenSignAlg"`

	// The duration of access tokens.
	AccessTokenDuration uint64 `json:"accessTokenDuration"`

	// The duration of refresh tokens.
	RefreshTokenDuration uint64 `json:"refreshTokenDuration"`

	// The duration of ID tokens.
	IdTokenDuration uint64 `json:"idTokenDuration"`

	// The duration of authorization response JWTs.
	AuthorizationResponseDuration uint64 `json:"authorizationResponseDuration"`

	// The key ID to identify a JWK used for signing JWT-based access tokens.
	AccessTokenSignatureKeyId string `json:"accessTokenSignatureKeyId"`

	// The key ID to identify a JWK used for signing authorization responses.
	AuthorizationSignatureKeyId string `json:"authorizationSignatureKeyId"`

	// The key ID to identify a JWK used for signing ID tokens.
	IdTokenSignatureKeyId string `json:"idTokenSignatureKeyId"`

	// The key ID to identify a JWK used for signing userinfo responses.
	UserInfoSignatureKeyId string `json:"userInfoSignatureKeyId"`

	// Supported backchannel token delivery modes.
	SupportedBackchannelTokenDeliveryModes []types.DeliveryMode `json:"supportedBackchannelTokenDeliveryModes"`

	// The URL of the backchannel authentication endpoint.
	BackchannelAuthenticationEndpoint string `json:"backchannelAuthenticationEndpoint"`

	// The flag which indicates whether `user_code` is supported at the backchannel authentication endpoint.
	BackchannelUserCodeParameterSupported bool `json:"backchannelUserCodeParameterSupported"`

	// The duration of backchannel authentication request IDs.
	BackchannelAuthReqIdDuration uint64 `json:"backchannelAuthReqIdDuration"`

	// The minimum interval in seconds between polling requests to the token endpoint in CIBA flows.
	BachcannelPollingInterval uint16 `json:"backchannelPollingInterval"`

	// The flag which indicates whether `binding_message` is always required for FAPI.
	BackchannelBindingMessageRequiredInFapi bool `json:"backchannelBindingMessageRequiredInFapi"`

	// The allowable clock skew in seconds between the server and clients.
	AllowableClockSkew uint16 `json:"allowableClockSkew"`

	// The URL of the device authorization endpoint.
	DeviceAuthorizationEndpoint string `json:"deviceAuthorizationEndpoint"`

	// The verification URI for the device flow.
	DeviceVerificationUri string `json:"deviceVerificationUri"`

	// The verification URI for the device flow with a placeholder for a user code.
	DeviceVerificationUriComplete string `json:"deviceVerificationUriComplete"`

	// The duration of device verification codes and end-user verification codes.
	DeviceFlowCodeDuration uint64 `json:"deviceFlowCodeDuration"`

	// The minimum interval in seconds between polling requests to the token endpoint in Device Flow.
	DeviceFlowPollingInterval uint16 `json:"deviceFlowPollingInterval"`

	// The character set for end-user verification codes for Device Flow.
	UserCodeCharset types.UserCodeCharset `json:"userCodeCharset"`

	// The length of end-user verification codes for Device Flow.
	UserCodeLength uint8 `json:"userCodeLength"`

	// The URL of the request object endpoint.
	RequestObjectEndpoint string `json:"requestObjectEndpoint"`

	// MTLS endpoint aliases.
	MtlsEndpointAliases []NamedUri `json:"mtlsEndpointAliases"`
}
