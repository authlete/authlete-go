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

type AuthorizationResponse struct {
	ApiResponse

	// The next action that the authorization server should take.
	Action AuthorizationAction `json:"action,omitempty"`

	//
	Service Service `json:"service,omitempty"`

	//
	Client Client `json:"client,omitempty"`

	//
	Display types.Display `json:"display,omitempty"`

	//
	MaxAge uint32 `json:"maxAge,omitempty"`

	//
	Scopes []Scope `json:"scopes,omitempty"`

	//
	DynamicScopes []DynamicScope `json:"dynamicScopes,omitempty"`

	//
	UiLocales []string `json:"uiLocales,omitempty"`

	//
	ClaimsLocales []string `json:"claimsLocales,omitempty"`

	//
	Claims []string `json:"claims,omitempty"`

	//
	AcrEssential bool `json:"acrEssential,omitempty"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed,omitempty"`

	//
	Acrs []string `json:"acrs,omitempty"`

	//
	Subject string `json:"subject,omitempty"`

	//
	LoginHint string `json:"loginHint,omitempty"`

	//
	Prompts []types.Prompt `json:"prompts,omitempty"`

	//
	RequestObjectPayload string `json:"requestObjectPayload,omitempty"`

	//
	IdTokenClaims string `json:"idTokenClaims,omitempty"`

	//
	UserInfoClaims string `json:"userInfoClaims,omitempty"`

	//
	Resources []string `json:"resources,omitempty"`

	//
	Purpose string `json:"purpose,omitempty"`

	//
	ResponseContent string `json:"responseContent,omitempty"`

	//
	Ticket string `json:"ticket,omitempty"`
}
