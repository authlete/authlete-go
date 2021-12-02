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

type BackchannelAuthenticationCompleteResponse struct {
	ApiResponse

	//
	Action BackchannelAuthenticationCompleteAction `json:"action,omitempty"`

	//
	ResponseContent string `json:"responseContent,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	ClientIdAlias string `json:"clientIdAlias,omitempty"`

	//
	ClientName string `json:"clientName,omitempty"`

	//
	DeliveryMode types.DeliveryMode `json:"deliveryMode,omitempty"`

	//
	ClientNotificationEndpoint string `json:"clientNotificationEndpoint,omitempty"`

	//
	ClientNotificationToken string `json:"clientNotificationToken,omitempty"`

	//
	AuthReqId string `json:"authReqId,omitempty"`

	//
	AccessToken string `json:"accessToken,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty"`

	//
	IdToken string `json:"idToken,omitempty"`

	//
	AccessTokenDuration uint64 `json:"accessTokenDuration,omitempty"`

	//
	RefreshTokenDuration uint64 `json:"refreshTokenDuration,omitempty"`

	//
	IdTokenDuration uint64 `json:"idTokenDuration,omitempty"`

	//
	JwtAccessToken string `json:"jwtAccessToken,omitempty"`

	//
	Resources []string `json:"resources,omitempty"`

	//
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`

	//
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
}
