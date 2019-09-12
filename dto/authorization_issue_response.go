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

// Response from Authlete's /api/auth/authorization/issue API.
type AuthorizationIssueResponse struct {
	ApiResponse

	// The next action that the authorization server should take.
	Action AuthorizationIssueAction `json:"action,omitempty"`

	// The response content which can be used to generated a response to the client.
	ResponseContent string `json:"responseContent,omitempty"`

	// The issued access token.
	//
	// An access token is issued when the `response_type` request parameter of
	// the authorization request includes `token`.
	AccessToken string `json:"accessToken,omitempty"`

	// The expiration date of the access token.
	//
	// The date in milliseconds since the Unix epoch at which the access token
	//  will expire.
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt,omitempty"`

	// The duration of the access token in seconds.
	AccessTokenDuration uint64 `json:"accessTokenDuration,omitempty"`

	// The issued ID token.
	//
	// An ID token is issued when the `response_type` request parameter of the
	// authorization request includes `id_token`.
	IdToken string `json:"idToken,omitempty"`

	// The issued authorization code.
	//
	// An authorization code is issued when the `response_type` request parameter
	// of the authorization request includes `code`.
	AuthorizationCode string `json:"authorizationCode,omitempty"`

	// The issued acces token in JWT format.
	//
	// If the authorization server is configued to issue JWT-based access tokens
	// (= if Service.AccessTokenSignAlg returns a non-nil value), a JWT-based
	// access token is issued along with the original random-string one.
	JwtAcessToken string `json:"jwtAccessToken,omitempty"`
}
