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

type Client struct {
	// The developer of this client.
	Developer string `json:"developer,omitempty"`

	// The client ID
	ClientId uint64 `json:"clientId,omitempty"`

	// The alias of the client ID.
	ClientIdAlias string `json:"clientIdAlias,omitempty"`

	// The flag which indicates whether the feature of Clien ID Alias is enabled.
	ClientIdAliasEnabled bool `json:"clientIdAliasEnabled,omitempty"`

	// The client secret.
	ClientSecret string `json:"clientSecret,omitempty"`

	// The client type.
	ClientType types.ClientType `json:"clientType,omitempty"`

	// Redirect URIs.
	RedirectUris []string `json:"redirectUris,omitempty"`

	// Response types that this client declares it may use.
	ResponseTypes []types.ResponseType `json:"responseTypes,omitempty"`

	// Grant types that this client declares it may use.
	GrantTypes []types.GrantType `json:"grantTypes,omitempty"`

	// The application type.
	ApplicationType types.ApplicationType `json:"applicationType,omitempty"`

	// Email addresses of contacts.
	Contacts []string `json:"contacts,omitempty"`

	// The name of the client.
	ClientName string `json:"clientName,omitempty"`

	// Client names for various locales.
	ClientNames []TaggedValue `json:"clientNames,omitempty"`

	// The URL where the logo image is located.
	LogoUri string `json:"logoUri,omitempty"`

	// Logo URIs for various locales.
	LogoUris []TaggedValue `json:"logoUris,omitempty"`

	// The URL of the website for the client.
	ClientUri string `json:"clientUri,omitempty"`

	// Client URIs for various locales.
	ClientUris []TaggedValue `json:"clientUris,omitempty"`

	// The URL of the policy page.
	PolicyUri string `json:"policyUri,omitempty"`

	// Policy URIs for various locales.
	PolicyUris []TaggedValue `json:"policyUris,omitempty"`

	// The URL of the Terms Of Service page.
	TosUri string `json:"tosUri,omitempty"`

	// TOS URIs for various locales.
	TosUris []TaggedValue `json:"tosUris,omitempty"`

	// The URL of the JWK Set document.
	JwksUri string `json:"jwksUri,omitempty"`

	// The JWK Set document.
	Jwks string `json:"jwks,omitempty"`

	// Calculated sector identifier host component.
	DerivedSectorIdentifier string `json:"derivedSectorIdentifier,omitempty"`

	// The sector identifier URI.
	SectorIdentifierUri string `json:"sectorIdentifierUri,omitempty"`

	// The subject type.
	SubjectType types.SubjectType `json:"subjectType,omitempty"`

	// JWS 'alg' for ID tokens.
	IdTokenSignAlg types.JWSAlg `json:"idTokenSignAlg,omitempty"`

	// JWE 'alg' for ID tokens.
	IdTokenEncryptionAlg types.JWEAlg `json:"idTokenEncryptionAlg,omitempty"`

	// JWE 'enc' for ID tokens.
	IdTokenEncryptionEnc types.JWEEnc `json:"idTokenEncryptionEnc,omitempty"`

	// JWS 'alg' for userinfo responses.
	UserInfoSignAlg types.JWSAlg `json:"userInfoSignAlg,omitempty"`

	// JWE 'alg' for userinfo responses.
	UserInfoEncryptionAlg types.JWEAlg `json:"userInfoEncryptionAlg,omitempty"`

	// JWE 'enc' for userinfo responses.
	UserInfoEncryptionEnc types.JWEEnc `json:"userInfoEncryptionEnc,omitempty"`

	// JWS 'alg' for request objects.
	RequestSignAlg types.JWSAlg `json:"requestSignAlg,omitempty"`

	// JWE 'alg' for request objects.
	RequestEncryptionAlg types.JWEAlg `json:"requestEncryptionAlg,omitempty"`

	// JWE 'enc' for request objects.
	RequestEncryptionEnc types.JWEEnc `json:"requestEncryptionEnc,omitempty"`

	// Client authentication method at the token endpoint.
	TokenAuthMethod types.ClientAuthMethod `json:"tokenAuthMethod,omitempty"`

	// JWS 'alg' for client assertions at the token endpoint.
	TokenAuthSignAlg types.JWSAlg `json:"tokenAuthSignAlg,omitempty"`

	// The default max age.
	DefaultMaxAge uint32 `json:"defaultMaxAge,omitempty"`

	// Default ACR values.
	DefaultAcrs []string `json:"defaultAcrs,omitempty"`

	// The flag which indicates whether this client always requires `auth_time`.
	AuthTimeRequired bool `json:"authTimeRequired,omitempty"`

	// The URL that can initiate login for this client application.
	LoginUri string `json:"loginUri,omitempty"`

	// The request URIs that this client declares it may use.
	RequestUris []string `json:"requestUris,omitempty"`

	// The description about this client.
	Description string `json:"description,omitempty"`

	// Descriptions for various locales.
	Descriptions []TaggedValue `json:"descriptions,omitempty"`

	// The time at which this client was created. Milliseconds since the Unix epoch.
	CreatedAt uint64 `json:"createdAt,omitempty"`

	// The time at which this client was last modified. MIlliseconds since the Unix epoch.
	ModifiedAt uint64 `json:"modifiedAt,omitempty"`

	// The extended information about this client.
	Extension ClientExtension `json:"extension,omitempty"`

	// The subject distinguished name of the certificate this client will use in MTLS.
	TlsClientAuthSubjectDn string `json:"tlsClientAuthSubjectDn,omitempty"`

	// The DNS subject alternative name of the certificate this client will use in MTLS.
	TlsClientAuthSanDns string `json:"tlsClientAuthSanDns,omitempty"`

	// The URI subject alternative name of the certificate this client will use in MTLS.
	TlsClientAuthSanUri string `json:"tlsClientAuthSanUri,omitempty"`

	// The IP address subject alternative name of the certificate this client will use in MTLS.
	TlsClientAuthSanIp string `json:"tlsClientAuthSanIp,omitempty"`

	// The email subject alternative name of the certificate this client will use in MTLS.
	TlsClientAuthSanEmail string `json:"tlsClientAuthSanEmail,omitempty"`

	// The flag which indicates whether certificate binding is enabled.
	TlsClientCertificateBoundAccessTokens bool `json:"tlsClientCertificateBoundAccessTokens,omitempty"`

	// The key ID of the JWK that represents a self-signed certificate used for client authentication.
	SelfSignedCertificateKeyId string `json:"selfSignedCertificateKeyId,omitempty"`

	// The software ID.
	SoftwareId string `json:"softwareId,omitempty"`

	// The software version
	SoftwareVersion string `json:"softwareVersion,omitempty"`

	// JWS 'alg' for authorization responses in JWT format (JARM).
	AuthorizationSignAlg types.JWSAlg `json:"authorizationSignAlg,omitempty"`

	// JWE 'alg' for authorization responses in JWT format (JARM).
	AuthorizationEncryptionAlg types.JWEAlg `json:"authorizationEncryptionAlg,omitempty"`

	// JWE 'enc' for authorization responses in JWT format (JARM).
	AuthorizationEncryptionEnc types.JWEEnc `json:"authorizationEncryptionEnc,omitempty"`

	// Backchannel token delivery mode.
	BcDeliveryMode types.DeliveryMode `json:"bcDeliveryMode,omitempty"`

	// Backchannel client notification endpoint.
	BcNotificationEndpoint string `json:"bcNotificationEndpoint,omitempty"`

	// JWS 'alg' for backchannel authentication request in JWT format.
	BcRequestSignAlg types.JWSAlg `json:"bcRequestSignAlg,omitempty"`

	// The flag which indicates whether user_code is required in backchannel authentication request.
	BcUserCodeRequired bool `json:"bcUserCodeRequired,omitempty"`

	// The flag which indicates whether this client has been registered dynamically.
	DynamicallyRegistered bool `json:"dynamicallyRegistered,omitempty"`

	// The hash of the registration access token.
	RegistrationAccessTokenHash string `json:"registrationAccessTokenHash,omitempty"`

	// The data types that this client may use as values of the `type` field
	// in `authorization_details`.
	AuthorizationDetailsTypes []string `json:"authorizationDetailsTypes,omitempty"`

	// The flag which indicates whether this client is required to use PAR
	// (OAuth 2.0 Pushed Authorization Requests).
	ParRequired bool `json:"parRequired,omitempty"`

	// The flag which indicates whether authorization requests from this client
	// are always required to utilize a request object by using either `request`
	// or `request_uri` request parameter.
	RequestObjectRequired bool `json:"requestObjectRequired,omitempty"`

	// Arbitrary attributes associated with this client.
	Attributes []Pair `json:"attributes,omitempty"`

	// Custom metadata supported by this client.
	CustomMetadata string `json:"customMetadata,omitempty"`

	// The flag which indicates whether encryption of request object is required
	// when the request object is passed through the front channel.
	FrontChannelRequestObjectEncryptionRequired bool `json:"frontChannelRequestObjectEncryptionRequired,omitempty"`

	// The flag which indicates whether the JWE alg of encrypted request
	// object must match the value of the request_object_encryption_alg client metadata.
	RequestObjectEncryptionAlgMatchRequired bool `json:"requestObjectEncryptionAlgMatchRequired,omitempty"`

	// The flag which indicates whether the JWE enc of encrypted request
	// object must match the value of the request_object_encryption_enc client metadata.
	RequestObjectEncryptionEncMatchRequired bool `json:"requestObjectEncryptionEncMatchRequired,omitempty"`
}
