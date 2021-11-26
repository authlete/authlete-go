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

type BackchannelAuthenticationResponse struct {
	ApiResponse

	//
	Action BackchannelAuthenticationAction `json:"action,omitempty"`

	//
	ResponseContent string `json:"responseContent,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	ClientIdAlias string `json:"clientIdAlias,omitempty"`

	//
	ClientName string `json:"clientName,omitempty"`

	//
	ClientAuthMethod types.ClientAuthMethod `json:"clientAuthMethod,omitempty"`

	//
	DeliveryMode types.DeliveryMode `json:"deliveryMode,omitempty"`

	//
	Scopes []Scope `json:"scopes,omitempty"`

	//
	DynamicScopes []DynamicScope `json:"dynamicScopes,omitempty"`

	//
	ClaimNames []string `json:"claimNames,omitempty"`

	//
	ClientNotificationToken string `json:"clientNotificationToken,omitempty"`

	//
	Acrs []string `json:"acrs,omitempty"`

	//
	HintType types.UserIdentificationHintType `json:"hintType,omitempty"`

	//
	Hint string `json:"hint,omitempty"`

	//
	Sub string `json:"sub,omitempty"`

	//
	BindingMessage string `json:"bindingMessage,omitempty"`

	//
	UserCode string `json:"userCode,omitempty"`

	//
	UserCodeRequired bool `json:"userCodeRequired,omitempty"`

	//
	RequestedExpiry uint32 `json:"requestedExpiry,omitempty"`

	//
	RequestContext string `json:"requestContext,omitempty"`

	//
	Resources []string `json:"resources,omitempty"`

	//
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`

	//
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`

	//
	Warnings []string `json:"warnings,omitempty"`

	//
	Ticket string `json:"ticket,omitempty"`
}
