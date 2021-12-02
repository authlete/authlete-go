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

type TokenIssueResponse struct {
	ApiResponse

	//
	Action TokenIssueAction `json:"action,omitempty"`

	//
	ResponseContent string `json:"responseContent,omitempty"`

	//
	AccessToken string `json:"accessToken,omitempty"`

	//
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt,omitempty"`

	//
	AccessTokenDuration uint64 `json:"accessTokenDuration,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty"`

	//
	RefreshTokenExpiresAt uint64 `json:"refreshTokenExpiresAt,omitempty"`

	//
	RefreshTokenDuration uint64 `json:"refreshTokenDuration,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	ClientIdAlias string `json:"clientIdAlias,omitempty"`

	//
	Subject string `json:"subject,omitempty"`

	//
	Scopes []string `json:"scopes,omitempty"`

	//
	Properties []Property `json:"properties,omitempty"`

	//
	JwtAccessToken string `json:"jwtAccessToken,omitempty"`

	//
	AccessTokenResources []string `json:"accessTokenResources,omitempty"`

	//
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`

	//
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
}
