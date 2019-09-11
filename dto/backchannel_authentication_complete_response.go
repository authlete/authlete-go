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

type BackchannelAuthenticationCompleteResponse struct {
	ApiResponse

	//
	Action BackchannelAuthenticationCompleteAction `json:"action"`

	//
	ResponseContent string `json:"responseContent"`

	//
	ClientId uint64 `json:"clientId"`

	//
	ClientIdAlias string `json:"clientIdAlias"`

	//
	ClientName string `json:"clientName"`

	//
	DeliveryMode types.DeliveryMode `json:"deliveryMode"`

	//
	ClientNotificationEndpoint string `json:"clientNotificationEndpoint"`

	//
	ClientNotificationToken string `json:"clientNotificationToken"`

	//
	AuthReqId string `json:"authReqId"`

	//
	AccessToken string `json:"accessToken"`

	//
	RefreshToken string `json:"refreshToken"`

	//
	IdToken string `json:"idToken"`

	//
	AccessTokenDuration uint64 `json:"accessTokenDuration"`

	//
	RefreshTokenDuration uint64 `json:"refreshTokenDuration"`

	//
	IdTokenDuration uint64 `json:"idTokenDuration"`

	//
	JwtAccessToken string `json:"jwtAccessToken"`
}
