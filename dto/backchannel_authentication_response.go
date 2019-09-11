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

type BackchannelAuthenticationResponse struct {
	ApiResponse

	//
	Action BackchannelAuthenticationAction `json:"action"`

	//
	ResponseContent string `json:"responseContent"`

	//
	ClientId uint64 `json:"clientId"`

	//
	ClientIdAlias string `json:"clientIdAlias"`

	//
	ClientName string `json:"clientName"`

	//
	ClientAuthMethod types.ClientAuthMethod `json:"clientAuthMethod"`

	//
	DeliveryMode types.DeliveryMode `json:"deliveryMode"`

	//
	Scopes []Scope `json:"scopes"`

	//
	ClaimNames []string `json:"claimNames"`

	//
	ClientNotificationToken string `json:"clientNotificationToken"`

	//
	Acrs []string `json:"acrs"`

	//
	HintType types.UserIdentificationHintType `json:"hintType"`

	//
	Hint string `json:"hint"`

	//
	Sub string `json:"sub"`

	//
	BindingMessage string `json:"bindingMessage"`

	//
	UserCode string `json:"userCode"`

	//
	UserCodeRequired bool `json:"userCodeRequired"`

	//
	RequestedExpiry uint32 `json:"requestedExpiry"`

	//
	RequestContext string `json:"requestContext"`

	//
	Warnings []string `json:"warnings"`

	//
	Ticket string `json:"ticket"`
}
