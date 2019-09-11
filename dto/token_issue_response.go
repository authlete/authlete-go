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

type TokenIssueResponse struct {
	ApiResponse

	//
	Action TokenIssueAction `json:"action"`

	//
	ResponseContent string `json:"responseContent"`

	//
	AccessToken string `json:"accessToken"`

	//
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt"`

	//
	AccessTokenDuration uint64 `json:"accessTokenDuration"`

	//
	RefreshToken string `json:"refreshToken"`

	//
	RefreshTokenExpiresAt uint64 `json:"refreshTokenExpiresAt"`

	//
	RefreshTokenDuration uint64 `json:"refreshTokenDuration"`

	//
	ClientId uint64 `json:"clientId"`

	//
	ClientIdAlias string `json:"clientIdAlias"`

	//
	Subject string `json:"subject"`

	//
	Scopes []string `json:"scopes"`

	//
	Properties []Property `json:"properties"`

	//
	JwtAccessToken string `json:"jwtAccessToken"`
}
