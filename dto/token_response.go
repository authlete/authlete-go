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

type TokenResponse struct {
	ApiResponse

	//
	Action TokenAction `json:"action"`

	//
	ResponseContent string `json:"responseContent"`

	//
	Username string `json:"username"`

	//
	Password string `json:"password"`

	//
	Ticket string `json:"ticket"`

	//
	AccessToken string `json:"accessToken"`

	//
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt"`

	//
	AccessTokenDuration uint64 `json:"accessTokenDuration"`

	//
	RefreshToken string `json:"refreshToken"`

	//
	RefresthTokenExpiresAt string `json:"refreshTokenExpiresAt"`

	//
	RefreshTokenDuration string `json:"refreshTokenDuration"`

	//
	IdToken string `json:"idToken"`

	//
	GrantType types.GrantType `json:"grantType"`

	//
	ClientId uint64 `json:"clientId"`

	//
	ClientIdAlias string `json:"clientIdAlias"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed"`

	//
	Subject string `json:"subject"`

	//
	Scopes []string `json:"scopes"`

	//
	Properties []Property `json:"properties"`

	//
	JwtAccessToken string `json:"jwtAccessToken"`

	//
	ClientAuthMethod types.ClientAuthMethod `json:"clientAuthMethod"`
}
