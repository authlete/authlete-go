//
// Copyright (C) 2019-2022 Authlete, Inc.
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

type TokenCreateRequest struct {
	//
	GrantType types.GrantType `json:"grantType,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	Subject string `json:"subject,omitempty"`

	//
	Scopes []string `json:"scopes,omitempty"`

	//
	AccessTokenDuration uint64 `json:"accessTokenDuration,omitempty"`

	//
	RefreshTokenDuration uint64 `json:"refreshTokenDuration,omitempty"`

	//
	Properties []Property `json:"properties,omitempty"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed,omitempty"`

	//
	AccessToken string `json:"accessToken,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty"`

	//
	AccessTokenPersistent bool `json:"accessTokenPersistent,omitempty"`

	//
	CertificateThumbprint string `json:"certificateThumbprint,omitempty"`

	//
	DpopKeyThumbprint string `json:"dpopKeyThumbprint,omitempty"`

	// Resources associated with the newly issued access token.
	//
	// Since v1.1.5.
	Resources []string `json:"resources,omitempty"`

	// Flag indicating whether the newly issued access token is for an
	// external attachment. See OpenID Connect for Identity Assurance
	// for details.
	//
	// Since v1.1.5.
	ForExternalAttachment bool `json:"forExternalAttachment,omitempty"`

	// Additional claims that are added to the payload part of the JWT
	// access token.
	//
	// Since v1.1.5.
	JwtAtClaims string `json:"jwtAtClaims,omitempty"`
}
