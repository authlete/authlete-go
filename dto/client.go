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

type Client struct {
	// The developer of this client.
	Developer string `json:"developer"`

	// The client ID
	ClientId uint64 `json:"clientId"`

	// The alias of the client ID.
	ClientIdAlias string `json:"clientIdAlias"`

	// The flag which indicates whether the feature of Clien ID Alias is enabled.
	ClientIdAliasEnabled bool `json:"clientIdAliasEnabled"`

	// The client secret.
	ClientSecret string `json:"clientSecret"`

	// The client type.
	ClientType types.ClientType `json:"clientType"`

	// Redirect URIs.
	RedirectUris []string `json:"redirectUris"`

	// Response types that this client declares it may use.
	ResponseTypes []types.ResponseType `json:"responseTypes"`

	// Grant types that this client declares it may use.
	GrantTypes []types.GrantType `json:"grantTypes"`

	// The application type.
	ApplicationType types.ApplicationType `json:"applicationType"`

	// Email addresses of contacts.
	Contacts []string `json:"contacts"`

	// The name of the client.
	ClientName string `json:"clientName"`

	// Client names for various locales.
	ClientNames []TaggedValue `json:"clientNames"`

	// The URL where the logo image is located.
	LogoUri string `json:"logoUri"`

	// Logo URIs for various locales.
	LogoUris []TaggedValue `json:"logoUris"`

	// The URL of the website for the client.
	ClientUri string `json:"clientUri"`

	// Client URIs for various locales.
	ClientUris []TaggedValue `json:"clientUris"`

	// The URL of the policy page.
	PolicyUri string `json:"policyUri"`

	// Policy URIs for various locales.
	PolicyUris []TaggedValue `json:"policyUris"`

	// The URL of the Terms Of Service page.
	TosUri string `json:"tosUri"`

	// TOS URIs for various locales.
	TosUris []TaggedValue `json:"tosUris"`

	// The URL of the JWK Set document.
	JwksUri string `json:"jwksUri"`

	// The sector identifier computed based on the sector identifier URI or redirect URIs.
	SectorIdentifier string `json:"sectorIdentifier"`

	// The sector identifier URI.
	SectorIdentifierUri string `json:"sectorIdentifierUri"`

	// The subject type.
	SubjectType types.SubjectType `json:"subjectType"`

	// JWS 'alg' for ID tokens.
	IdTokenSignAlg types.JWSAlg `json:"idTokenSignAlg"`

	// JWE 'alg' for ID tokens.
	IdTokenEncryptionAlg types.JWEAlg `json:"idTokenEncryptionAlg"`

	// JWE 'enc' for ID tokens.
	IdTokenEncryptionEnc types.JWEEnc `json:"idTokenEncryptionEnc"`

	// JWS 'alg' for userinfo responses.
	UserInfoSignAlg types.JWSAlg `json:"userInfoSignAlg"`

	// JWE 'alg' for userinfo responses.
	UserInfoEncryptionAlg types.JWEAlg `json:"userInfoEncryptionAlg"`

	// JWE 'enc' for userinfo responses.
	UserInfoEncryptionEnc types.JWEEnc `json:"userInfoEncryptionEnc"`

	// JWS 'alg' for request objects.
	RequestSignAlg types.JWSAlg `json:"requestSignAlg"`

	// JWE 'alg' for request objects.
	RequestEncryptionAlg types.JWEAlg `json:"requestEncryptionAlg"`

	// JWE 'enc' for request objects.
	RequestEncryptionEnc types.JWEEnc `json:"requestEncryptionEnc"`

	// Client authentication method at the token endpoint.
	TokenAuthMethod types.ClientAuthMethod `json:"tokenAuthMethod"`

	// JWS 'alg' for client assertions at the token endpoint.
	TokenAuthSignAlg types.JWSAlg `json:"tokenAuthSignAlg"`

	// The default max age.
	DefaultMaxAge uint32 `json:"defaultMaxAge"`

	// Default ACR values.
	DefaultAcrs []string `json:"defaultAcrs"`

	// The flag which indicates whether this client always requires `auth_time`.
	AuthTimeRequired bool `json:"authTimeRequired"`

	// The URL that can initiate login for this client application.
	LoginUri string `json:"loginUri"`

	// The request URIs that this client declares it may use.
	RequestUris []string `json:"requestUri"`

	// The description about this client.
	Description string `json:"description"`

	// Descriptions for various locales.
	Descriptions []TaggedValue `json:"descriptions"`

	// The time at which this client was created. Milliseconds since the Unix epoch.
	CreatedAt uint64 `json:"createdAt"`

	// The time at which this client was last modified. MIlliseconds since the Unix epoch.
	ModifiedAt uint64 `json:"modifiedAt"`

	// The extended information about this client.
	Extension ClientExtension `json:"extension"`

	// The subject distinguished name of the certificate this client will use in MTLS.
	TlsClientAuthSubjectDn string `json:"tlsClientAuthSubjectDn"`

	// The DNS subject alternative name of the certificate this client will use in MTLS.
	TlsClientAuthSanDns string `json:"tlsClientAuthSanDns"`

	// The URI subject alternative name of the certificate this client will use in MTLS.
	TlsClientAuthSanUri string `json:"tlsClientAuthSanUri"`

	// The IP address subject alternative name of the certificate this client will use in MTLS.
	TlsClientAuthSanIp string `json:"tlsClientAuthSanIp"`

	// The email subject alternative name of the certificate this client will use in MTLS.
	TlsClientAuthSanEmail string `json:"tlsClientAuthSanEmail"`

	// The flag which indicates whether certificate binding is enabled.
	TlsClientCertificateBoundAccessTokens bool `json:"tlsClientCertificateBoundAccessTokens"`

	// The key ID of the JWK that represents a self-signed certificate used for client authentication.
	SelfSignedCertificateKeyId string `json:"selfSignedCertificateKeyId"`

	// The software ID.
	SoftwareId string `json:"softwareId"`

	// The software version
	SoftwareVersion string `json:"softwareVersion"`

	// JWS 'alg' for authorization responses in JWT format (JARM).
	AuthorizationSignAlg types.JWSAlg `json:"authorizationSignAlg"`

	// JWE 'alg' for authorization responses in JWT format (JARM).
	AuthorizationEncryptionAlg types.JWEAlg `json:"authorizationEncryptionAlg"`

	// JWE 'enc' for authorization responses in JWT format (JARM).
	AuthorizationEncryptionEnc types.JWEEnc `json:"authorizationEncryptionEnc"`

	// Backchannel token delivery mode.
	BcDeliveryMode types.DeliveryMode `json:"bcDeliveryMode"`

	// Backchannel client notification endpoint.
	BcNotificationEndpoint string `json:"bcNotificationEndpoint"`

	// JWS 'alg' for backchannel authentication request in JWT format.
	BcRequestSignAlg types.JWSAlg `json:"bcRequestSignAlg"`

	// The flag which indicates whether user_code is required in backchannel authentication request.
	BcUserCodeRequired bool `json:"bcUserCodeRequired"`

	// The flag which indicates whether this client has been registered dynamically.
	DynamicallyRegistered bool `json:"dynamicallyRegistered"`

	// The hash of the registration access token.
	RegistrationAccessTokenHash string `json:"registrationAccessTokenHash"`
}
