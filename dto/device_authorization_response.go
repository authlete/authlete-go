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

type DeviceAuthorizationResponse struct {
	ApiResponse

	//
	Action DeviceAuthorizationAction `json:"action"`

	//
	ResponseContent string `json:"responseContet"`

	//
	ClientId uint64 `json:"clientId"`

	//
	ClientIdAlias string `json:"clientIdAlias"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed"`

	//
	ClientName string `json:"clientName"`

	//
	ClientAuthMethod types.ClientAuthMethod `json:"clientAuthMethod"`

	//
	Scopes []Scope `json:"scopes"`

	//
	ClaimNames []string `json:"claimNames"`

	//
	Acrs []string `json:"acrs"`

	//
	DeviceCode string `json:"deviceCode"`

	//
	UserCode string `json:"userCode"`

	//
	VerificationUri string `json:"verificationUri"`

	//
	VerificatinoUriComplete string `json:"verificationUriComplete"`

	//
	ExpiresIn uint32 `json:"expiresIn"`

	//
	Interval uint32 `json:"interval"`

	//
	Warnings []string `json:"warnings"`
}
