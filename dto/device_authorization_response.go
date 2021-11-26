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

type DeviceAuthorizationResponse struct {
	ApiResponse

	//
	Action DeviceAuthorizationAction `json:"action,omitempty"`

	//
	ResponseContent string `json:"responseContent,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	ClientIdAlias string `json:"clientIdAlias,omitempty"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed,omitempty"`

	//
	ClientName string `json:"clientName,omitempty"`

	//
	ClientAuthMethod types.ClientAuthMethod `json:"clientAuthMethod,omitempty"`

	//
	Scopes []Scope `json:"scopes,omitempty"`

	//
	DynamicScopes []DynamicScope `json:"dynamicScopes,omitempty"`

	//
	ClaimNames []string `json:"claimNames,omitempty"`

	//
	Acrs []string `json:"acrs,omitempty"`

	//
	DeviceCode string `json:"deviceCode,omitempty"`

	//
	UserCode string `json:"userCode,omitempty"`

	//
	VerificationUri string `json:"verificationUri,omitempty"`

	//
	VerificationUriComplete string `json:"verificationUriComplete,omitempty"`

	//
	ExpiresIn uint32 `json:"expiresIn,omitempty"`

	//
	Interval uint32 `json:"interval,omitempty"`

	//
	Resources []string `json:"resources,omitempty"`

	//
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`

	//
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`

	//
	Warnings []string `json:"warnings,omitempty"`
}
