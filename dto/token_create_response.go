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

type TokenCreateResponse struct {
	ApiResponse

	//
	Action TokenCreateAction `json:"action,omitempty"`

	//
	GrantType types.GrantType `json:"grantType,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	Subject string `json:"subject,omitempty"`

	//
	Scopes []string `json:"scopes,omitempty"`

	//
	AccessToken string `json:"accessToken,omitempty"`

	//
	TokenType string `json:"tokenType,omitempty"`

	//
	ExpiresIn uint64 `json:"expiresIn,omitempty"`

	//
	ExpiresAt uint64 `json:"expiresAt,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty"`

	//
	Properties []Property `json:"properties,omitempty"`

	// Access token in the JWT format.
	//
	// Since v1.5.0.
	JwtAccessToken string `json:"jwtAccessToken,omitempty"`

	// Flag indicating whether the access token is for an external attachment.
	// See OpenID Connect for Identity Assurance for details.
	//
	// Since v1.5.0.
	ForExternalAttachment bool `json:"forExternalAttachment,omitempty"`

	// Token ID.
	//
	// Since v1.5.0.
	TokenId string `json:"tokenId,omitempty"`
}
