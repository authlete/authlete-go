//
// Copyright (C) 2019-2023 Authlete, Inc.
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

type DeviceCompleteRequest struct {
	//
	UserCode string `json:"userCode,omitempty"`

	//
	Result DeviceCompleteResult `json:"result,omitempty"`

	//
	Sujbect string `json:"subject,omitempty"`

	//
	Sub string `json:"sub,omitempty"`

	//
	AuthTime uint64 `json:"authTime,omitempty"`

	//
	Acr string `json:"acr,omitempty"`

	//
	Claims string `json:"claims,omitempty"`

	//
	Properties []Property `json:"properties,omitempty"`

	//
	Scopes []string `json:"scopes,omitempty"`

	//
	IdtHeaderParams string `json:"idtHeaderParams,omitempty"`

	// Claims that the user has consented for the client application to know.
	//
	// Since v1.1.5.
	ConsentedClaims []string `json:"consentedClaims,omitempty"`

	// Additional claims that are added to the payload part of the JWT
	// access token.
	//
	// Since v1.1.5.
	JwtAtClaims string `json:"jwtAtClaims,omitempty"`

	//
	ErrorDescription string `json:"errorDescription,omitempty"`

	//
	ErrorUri string `json:"errorUri,omitempty"`

	// aud claim type of id token.
	//
	// Since v2.3.3.
	IdTokenAudType types.IdTokenAudType `json:"idTokenAudType,omitempty"`
}
