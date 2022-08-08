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

type TokenResponse struct {
	ApiResponse

	//
	Action TokenAction `json:"action,omitempty"`

	//
	ResponseContent string `json:"responseContent,omitempty"`

	//
	Username string `json:"username,omitempty"`

	//
	Password string `json:"password,omitempty"`

	//
	Ticket string `json:"ticket,omitempty"`

	//
	AccessToken string `json:"accessToken,omitempty"`

	//
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt,omitempty"`

	//
	AccessTokenDuration uint64 `json:"accessTokenDuration,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty"`

	//
	RefresthTokenExpiresAt uint64 `json:"refreshTokenExpiresAt,omitempty"`

	//
	RefreshTokenDuration uint64 `json:"refreshTokenDuration,omitempty"`

	//
	IdToken string `json:"idToken,omitempty"`

	//
	GrantType types.GrantType `json:"grantType,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	ClientIdAlias string `json:"clientIdAlias,omitempty"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed,omitempty"`

	//
	Subject string `json:"subject,omitempty"`

	//
	Scopes []string `json:"scopes,omitempty"`

	//
	Properties []Property `json:"properties,omitempty"`

	//
	JwtAccessToken string `json:"jwtAccessToken,omitempty"`

	//
	ClientAuthMethod types.ClientAuthMethod `json:"clientAuthMethod,omitempty"`

	//
	Resources []string `json:"resources,omitempty"`

	//
	AccessTokenResources []string `json:"accessTokenResources,omitempty"`

	// Grant ID. See Grant Management for OAuth 2.0 for details.
	//
	// Since v1.5.0.
	GrantId string `json:"grantId,omitempty"`

	//
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`

	//
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`

	// The values of the "audience" request parameters that are contained in
	// the token exchange request (RFC 8693).
	//
	// Since v1.5.0.
	Audiences []string `json:"audiences,omitempty"`

	// The value of the "requested_token_type" request parameter of the token
	// exchange request (RFC 8693).
	//
	// Since v1.5.0.
	RequestedTokenType types.TokenType `json:"requestedTokenType,omitempty"`

	// The value of the "subject_token" request parameter of the token exchange
	// request (RFC 8693).
	//
	// Since v1.5.0.
	SubjectToken string `json:"subjectToken,omitempty"`

	// The value of the "subject_token_type" request parameter of the token
	// exchange request (RFC 8693).
	//
	// Since v1.5.0.
	SubjectTokenType types.TokenType `json:"subjectTokenType,omitempty"`

	// The information about the token specified by the "subject_token" request
	// parameter. The information is available only when the token type is
	// either "urn:ietf:params:oauth:token-type:access_token" or
	// "urn:ietf:params:oauth:token-type:refresh_token".
	//
	// Since v1.5.0.
	SubjectTokenInfo TokenInfo `json:"subjectTokenInfo,omitempty"`

	// The value of the "actor_token" request parameter of the token exchange
	// request (RFC 8693).
	//
	// Since v1.5.0.
	ActorToken string `json:"actorToken,omitempty"`

	// The value of the "actor_token_type" request parameter of the token
	// exchange request (RFC 8693).
	//
	// Since v1.5.0.
	ActorTokenType types.TokenType `json:"actorTokenType,omitempty"`

	// The information about the token specified by the "actor_token" request
	// parameter. The information is available only when the token type is
	// either "urn:ietf:params:oauth:token-type:access_token" or
	// "urn:ietf:params:oauth:token-type:refresh_token".
	//
	// Since v1.5.0.
	ActorTokenInfo TokenInfo `json:"actorTokenInfo,omitempty"`
}
