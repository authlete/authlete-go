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

package dto

import "github.com/authlete/authlete-go/types"

type Service struct {
	// The name of the service.
	ServiceName string `json:"serviceName,omitempty"`

	// The API key of the service.
	ApiKey uint64 `json:"apiKey,omitempty"`

	// The API secret of the service.
	ApiSecret string `json:"apiSecret,omitempty"`

	// The issuer identifier.
	Issuer string `json:"issuer,omitempty"`

	// The URL of the authorization endpoint.
	AuthorizationEndpoint string `json:"authorizationEndpoint,omitempty"`

	// The URL of the token endpoint.
	TokenEndpoint string `json:"tokenEndpoint,omitempty"`

	// The URL of the revocation endpoint.
	RevocationEndpoint string `json:"revocationEndpoint,omitempty"`

	// Client authentication methods supported at the revocation endpoint.
	SupportedRevocationAuthMethods []types.ClientAuthMethod `json:"supportedRevocationAuthMethods,omitempty"`

	// The URL of the userinfo endpoint.
	UserInfoEndpoint string `json:"userInfoEndpoint,omitempty"`

	// The URI of the JWK Set document.
	JwksUri string `json:"jwksUri,omitempty"`

	// The content of the JWK Set document.
	Jwks string `json:"jwks,omitempty"`

	// The URL of the registration endpoint.
	RegistrationEndpoint string `json:"registrationEndpoint,omitempty"`

	// The URL of the registration management endpoint.
	RegistrationManagementEndpoint string `json:"registrationManagementEndpoint,omitempty"`

	// Supported scopes.
	SupportedScopes []Scope `json:"supportedScopes,omitempty"`

	// Supported response types.
	SupportedResponseTypes []types.ResponseType `json:"supportedResponseTypes,omitempty"`

	// Supported grant types.
	SupportedGrantTypes []types.GrantType `json:"supportedGrantTypes,omitempty"`

	// String supported ACRs.
	SupportedAcrs []string `json:"supportedAcrs,omitempty"`

	// Client authentication methods supported at the token endpoint.
	SupportedTokenAuthMethods []types.ClientAuthMethod `json:"supportedTokenAuthMethods,omitempty"`

	// Supported displays.
	SupportedDisplays []types.Display `json:"supportedDisplays,omitempty"`

	// Supported claim types.
	SupportedClaimTypes []types.ClaimType `json:"supportedClaimTypes,omitempty"`

	// Supported claims.
	SupportedClaims []string `json:"supportedClaims,omitempty"`

	// The URL of the service documentation.
	ServiceDocumentation string `json:"serviceDocumentation,omitempty"`

	// Supported claim locales.
	SupportedClaimLocales []string `json:"supportedClaimLocales,omitempty"`

	// Supported UI locales.
	SupportedUiLocales []string `json:"supportedUiLocales,omitempty"`

	// The URL of the policy document.
	PolicyUri string `json:"policyUri,omitempty"`

	// The URL of the terms of service.
	TosUri string `json:"tosUri,omitempty"`

	// The URL of the authentication callback endpoint.
	AuthenticationCallbackEndpoint string `json:"authenticationCallbackEndpoint,omitempty"`

	// The API key to access the authentication callback endpoint.
	AuthenticationCallbackApiKey string `json:"authenticationCallbackApiKey,omitempty"`

	// The API secret to access the authentication callback endpoint.
	AuthenticationCallbackApiSecret string `json:"authenticationCallbackApiSecret,omitempty"`

	// Supported SNSes.
	SupportedSnses []types.Sns `json:"supportedSnses,omitempty"`

	// Credentials of supported SNSes.
	SnsCredentials []SnsCredentials `json:"snsCredentials,omitempty"`

	// The URL of the developer authentication callback endpoint.
	DeveloperAuthenticationCallbackEndpoint string `json:"developerAuthenticationCallbackEndpoint,omitempty"`

	// The API key to access the developer authentication callback endpoint.
	DeveloperAuthenticationCallbackApiKey string `json:"developerAuthenticationCallbackApiKey,omitempty"`

	// The API secret to access the developer authentication callback endpoint.
	DeveloperAuthenticationCallbackApiSecret string `json:"developerAuthenticationCallbackApiSecret,omitempty"`

	// The upper limit of the number of client applications per developer.
	ClientsPerDeveloper uint16 `json:"clientsPerDeveloper,omitempty"`

	// The flag which indicates whether the direct authorization endpoint is enabled.
	DirectAuthorizationEndpointEnabled bool `json:"directAuthorizationEndpointEnabled,omitempty"`

	// The flag which indicates whether the direct token endpoint is enabled.
	DirectTokenEndpointEnabled bool `json:"directTokenEndpointEnabled,omitempty"`

	// The flag which indicates whether the direct revocation endpoint is enabled.
	DirectRevocationEndpointEnabled bool `json:"directRevocationEndpointEnabled,omitempty"`

	// The flag which indicates whether the direct userinfo endpoint is enabled.
	DirectUserInfoEndpointEnabled bool `json:"directUserInfoEndpointEnabled,omitempty"`

	// The flag which indicates whether the direct jwks endpoint is enabled.
	DirectJwksEndpointEnabled bool `json:"directJwksEndpointEnabled,omitempty"`

	// The flag which indicates whether the direct introspection endpoint is enabled.
	DirectIntrospectionEndpointEnabled bool `json:"directIntrospectionEndpointEnabled,omitempty"`

	// The flag which indicates whether the number of access tokens per subject is limited to at most 1.
	SingleAccessTokenPerSubject bool `json:"singleAccessTokenPerSubject,omitempty"`

	// The flag which indicates whether PKCE is always required.
	PkceRequired bool `json:"pkceRequired,omitempty"`

	// The flag which indicates whether S256 is always required for code_challenge_method.
	PkceS256Required bool `json:"pkceS256Required,omitempty"`

	// The flag which indicates whether refresh tokens remain valid after use.
	RefreshTokenKept bool `json:"refreshTokenKept,omitempty"`

	// The flag which indicates whether the remaining duration of the used
	// refresh token is taken over to the newly issued refresh token.
	RefreshTokenDurationKept bool `json:"refreshTokenDurationKept,omitempty"`

	// The flag which indicates whether duration of refresh tokens are
	// reset when they are used even if the refreshTokenKept property
	RefreshTokenDurationReset bool `json:"refreshTokenDurationReset,omitempty"`

	// The flag which indicates whether `error_description` is omitted.
	ErrorDescriptionOmitted bool `json:"errorDescriptionOmitted,omitempty"`

	// The flag which indicates whether `error_uri` is omitted.
	ErrorUriOmitted bool `json:"errorUriOmitted,omitempty"`

	// The flag which indicates whether the feature of Client ID Alias is enabled.
	ClientIdAliaseEnabled bool `json:"clientIdAliasEnabled,omitempty"`

	// Supported service profiles.
	SupportedServiceProfiles []types.ServiceProfile `json:"supportedServiceProfiles,omitempty"`

	// The flag which indicates whether certificate binding is supported.
	TlsClientCertificateBoundAccessTokens bool `json:"tlsClientCertificateBoundAccessTokens,omitempty"`

	// The URL of the introspection endpoint.
	IntrospectionEndpoint string `json:"introspectionEndpoint,omitempty"`

	// Client authentication methods supports at the introspection endpoint.
	SupportedIntrospectionAuthMethods []types.ClientAuthMethod `json:"supportedIntrospectionAuthMethods,omitempty"`

	// The flag which indicates whether certification chain is validated for MTLS.
	MutualTlsValidatePkiCertChain bool `json:"mutualTlsValidatePkiCertChain,omitempty"`

	// The list of trusted root certificates for MTLS.
	TrustedRootCertificates []string `json:"trustedRootCertificates,omitempty"`

	// The flag which indicates whether Dynamic Client Registration is supported.
	DynamicRegistrationSupported bool `json:"dynamicRegistrationSupported,omitempty"`

	// The URL of the end session endpoint.
	EndSessionEndpoint string `json:"endSessionEndpoint,omitempty"`

	// The description of this service.
	Description string `json:"description,omitempty"`

	// The type of access token.
	AccessTokenType string `json:"accessTokenType,omitempty"`

	// Signature algorithm of JWT-based access tokens.
	//
	// When this property is not nil, access tokens issued by this service are
	// JWS. Otherwise, access tokens are random strings as before.
	AccessTokenSignAlg types.JWSAlg `json:"accessTokenSignAlg,omitempty"`

	// The duration of access tokens.
	AccessTokenDuration uint64 `json:"accessTokenDuration,omitempty"`

	// The duration of refresh tokens.
	RefreshTokenDuration uint64 `json:"refreshTokenDuration,omitempty"`

	// The duration of ID tokens.
	IdTokenDuration uint64 `json:"idTokenDuration,omitempty"`

	// The duration of authorization response JWTs.
	AuthorizationResponseDuration uint64 `json:"authorizationResponseDuration,omitempty"`

	// The duration of pushed authorization requests.
	PushedAuthReqDuration uint64 `json:"pushedAuthReqDuration,omitempty"`

	// The key ID to identify a JWK used for signing JWT-based access tokens.
	AccessTokenSignatureKeyId string `json:"accessTokenSignatureKeyId,omitempty"`

	// The key ID to identify a JWK used for signing authorization responses.
	AuthorizationSignatureKeyId string `json:"authorizationSignatureKeyId,omitempty"`

	// The key ID to identify a JWK used for signing ID tokens.
	IdTokenSignatureKeyId string `json:"idTokenSignatureKeyId,omitempty"`

	// The key ID to identify a JWK used for signing userinfo responses.
	UserInfoSignatureKeyId string `json:"userInfoSignatureKeyId,omitempty"`

	// Supported backchannel token delivery modes.
	SupportedBackchannelTokenDeliveryModes []types.DeliveryMode `json:"supportedBackchannelTokenDeliveryModes,omitempty"`

	// The URL of the backchannel authentication endpoint.
	BackchannelAuthenticationEndpoint string `json:"backchannelAuthenticationEndpoint,omitempty"`

	// The flag which indicates whether `user_code` is supported at the backchannel authentication endpoint.
	BackchannelUserCodeParameterSupported bool `json:"backchannelUserCodeParameterSupported,omitempty"`

	// The duration of backchannel authentication request IDs.
	BackchannelAuthReqIdDuration uint64 `json:"backchannelAuthReqIdDuration,omitempty"`

	// The minimum interval in seconds between polling requests to the token endpoint in CIBA flows.
	BachcannelPollingInterval uint16 `json:"backchannelPollingInterval,omitempty"`

	// The flag which indicates whether `binding_message` is always required for FAPI.
	BackchannelBindingMessageRequiredInFapi bool `json:"backchannelBindingMessageRequiredInFapi,omitempty"`

	// The allowable clock skew in seconds between the server and clients.
	AllowableClockSkew uint16 `json:"allowableClockSkew,omitempty"`

	// The URL of the device authorization endpoint.
	DeviceAuthorizationEndpoint string `json:"deviceAuthorizationEndpoint,omitempty"`

	// The verification URI for the device flow.
	DeviceVerificationUri string `json:"deviceVerificationUri,omitempty"`

	// The verification URI for the device flow with a placeholder for a user code.
	DeviceVerificationUriComplete string `json:"deviceVerificationUriComplete,omitempty"`

	// The duration of device verification codes and end-user verification codes.
	DeviceFlowCodeDuration uint64 `json:"deviceFlowCodeDuration,omitempty"`

	// The minimum interval in seconds between polling requests to the token endpoint in Device Flow.
	DeviceFlowPollingInterval uint16 `json:"deviceFlowPollingInterval,omitempty"`

	// The character set for end-user verification codes for Device Flow.
	UserCodeCharset types.UserCodeCharset `json:"userCodeCharset,omitempty"`

	// The length of end-user verification codes for Device Flow.
	UserCodeLength uint8 `json:"userCodeLength,omitempty"`

	// The URL of the pushed authorization request endpoint.
	PushedAuthReqEndpoint string `json:"pushedAuthReqEndpoint,omitempty"`

	// MTLS endpoint aliases.
	MtlsEndpointAliases []NamedUri `json:"mtlsEndpointAliases,omitempty"`

	// Supported data types for authorization_details.
	SupportedAuthorizationDetailsTypes []string `json:"supportedAuthorizationDetailsTypes,omitempty"`

	// Supported trust frameworks. This corresponds to "trust_frameworks_supported".
	SupportedTrustFrameworks []string `json:"supportedTrustFrameworks,omitempty"`

	// Supported evidence. This corresponds to "evidence_supported".
	SupportedEvidence []string `json:"supportedEvidence,omitempty"`

	// Supported ID documents. This corresponds to "id_documents_supported".
	SupportedIdentityDocuments []string `json:"supportedDocuments,omitempty"`

	// Supported verification methods. This corresponds to "id_documents_verification_methods_supported".
	SupportedVerificationMethods []string `json:"supportedVerificationMethods,omitempty"`

	// Supported verified claims. This corresponds to "claims_in_verified_claims_supported".
	SupportedVerifiedClaims []string `json:"supportedVerifiedClaims,omitempty"`

	// The flag which indicates whether token requests from public client without
	// the `client_id` request parameter are allowed when the client can be guessed
	// from `authorization_code` or `refresh_token`.
	MissingClientIdAllowed bool `json:"missingClientIdAllowed,omitempty"`

	// The flag which indicates whether this service requires that clients use
	// PAR (OAuth 2.0 Pushed Authorization Requests).
	ParRequired bool `json:"parRequired,omitempty"`

	// The flag which indicates whether authorization requests must utilize a request object.
	RequestObjectRequired bool `json:"requestObjectRequired,omitempty"`

	// The flag which indicates traditional request object processing
	// (rules defined in OIDC Core 1.0) is applied.
	TraditionalRequestObjectProcessingApplied bool `json:"traditionalRequestObjectProcessingApplied,omitempty"`

	// The flag which indicates whether claims specified by shortcut scopes
	// (e.g. profile) are included in the issued ID token only when no access
	// token is issued.
	ClaimShortcutRestrictive bool `json:"claimShortcutRestrictive,omitempty"`

	// The flag which indicates whether requests that request no scope are
	// rejected or not.
	ScopeRequired bool `json:"scopeRequired,omitempty"`

	// The flag which indicates whether the nbf claim in the request
	// object is optional even when the authorization request is regarded
	// as a FAPI-Part2 request.
	NbfOptional bool `json:"nbfOptional,omitempty"`

	// The flag which indicates whether generation of the iss response
	// parameter is suppressed.
	IssSuppressed bool `json:"issSuppressed,omitempty"`

	// Arbitrary attributes associated with this service.
	Attributes []Pair `json:"attributes,omitempty"`

	// Custom client metadata supported by this service.
	SupportedCustomClientMetadata []string `json:"supportedCustomClientMetadata,omitempty"`

	// The flag which indicates whether the expiration date of an access token
	// never exceeds that of the corresponding refresh token.
	TokenExpirationLinked bool `json:"tokenExpirationLinked,omitempty"`

	// The flag which indicates whether encryption of request object is required
	// when the request object is passed through the front channel.
	FrontChannelRequestObjectEncryptionRequired bool `json:"frontChannelRequestObjectEncryptionRequired,omitempty"`

	// The flag which indicates whether the JWE alg of encrypted request
	// object must match the value of the request_object_encryption_alg client metadata.
	RequestObjectEncryptionAlgMatchRequired bool `json:"requestObjectEncryptionAlgMatchRequired,omitempty"`

	// The flag which indicates whether the JWE enc of encrypted request
	// object must match the value of the request_object_encryption_enc client metadata.
	RequestObjectEncryptionEncMatchRequired bool `json:"requestObjectEncryptionEncMatchRequired,omitempty"`

	// The flag which indicates whether HSM (Hardware Security Module) support is
	// enabled for this service.
	HsmEnabled bool `json:"hsmEnabled,omitempty"`

	// Hardware-secured keys. Output only.
	Hsks []Hsk `json:"hsks,omitempty"`

	// if the scope provided by developer on DCR/management can be requestable
	DcrScopeUsedAsRequestable bool `json:"dcrScopeUsedAsRequestable,omitempty"`

	LoopbackRedirectionUriVariable bool `json:"loopbackRedirectionUriVariable,omitempty"`
}
