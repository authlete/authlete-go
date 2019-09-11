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

type AuthorizationResponse struct {
	ApiResponse

	// The next action that the authorization server should take.
	Action AuthorizationAction `json:"action"`

	//
	Service Service `json:"service"`

	//
	Client Client `json:"client"`

	//
	Display types.Display `json:"display"`

	//
	MaxAge uint32 `json:"maxAge"`

	//
	Scopes []Scope `json:"scopes"`

	//
	UiLocales []string `json:"uiLocales"`

	//
	ClaimsLocales []string `json:"claimsLocales"`

	//
	Claims []string `json:"claims"`

	//
	AcrEssential bool `json:"acrEssential"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed"`

	//
	Acrs []string `json:"acrs"`

	//
	Subject string `json:"subject"`

	//
	LoginHint string `json:"loginHint"`

	//
	Prompts []types.Prompt `json:"prompts"`

	//
	RequestObjectPayload string `json:"requestObjectPayload"`

	//
	IdTokenClaims string `json:"idTokenClaims"`

	//
	UserInfoClaims string `json:"userInfoClaims"`

	//
	ResponseContent string `json:"responseContent"`

	//
	Ticket string `json:"ticket"`
}
