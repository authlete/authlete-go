//
// Copyright (C) 2019-2022 Authlete, Inc.
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

// Request to Authlete's /api/auth/authorization/issue API.
type AuthorizationIssueRequest struct {
	// The ticket issued by Authlete's /api/auth/authorization API.
	Ticket string `json:"ticket,omitempty"`

	// The subject (unique identifier) of the user.
	Subject string `json:"subject,omitempty"`

	// The value of the `sub` claim in an ID token.
	//
	// When this property is empty, the value of Subject is used as
	// the value of the `sub` claim.
	Sub string `json:"sub,omitempty"`

	// The time when the use was authenticated. (Seconds since epoch)
	AuthTime uint64 `json:"authTime,omitempty"`

	// The authentication context class reference.
	Acr string `json:"acr,omitempty"`

	// Claims in JSON format.
	Claims string `json:"claims,omitempty"`

	// Properties to associate with tokens.
	Properties []Property `json:"properties,omitempty"`

	// Scopes to associate with tokens.
	//
	// When this property is empty, the scopes specified in the original
	// authorization request are used.
	Scopes []string `json:"scopes,omitempty"`

	// JSON that represents additional JWS header parameters for ID tokens
	// that may be issued based on the authorization request.
	//
	// Since v1.1.0.
	IdtHeaderParams string `json:"idtHeaderParams,omitempty"`

	// Claims that the user has consented for the client application to know.
	//
	// Since v1.5.0.
	ConsentedClaims []string `json:"consentedClaims,omitempty"`

	// Claim key-value pairs that are used to compute values of transformed
	// claims. The format is JSON.
	//
	// Since v1.5.0.
	ClaimsForTx string `json:"claimsForTx,omitempty"`

	// Verified claim key-value pairs that are used to compute values of
	// transformed claims. The format of each element is JSON.
	//
	// Since v1.5.0.
	VerifiedClaimsForTx []string `json:"verifiedClaimsForTx,omitempty"`

	// Additional claims that are added to the payload part of the JWT
	// access token.
	//
	// Since v1.5.0.
	JwtAtClaims string `json:"jwtAtClaims,omitempty"`

	// The representation of an access token that may be issued as a result
	// of the Authlete API call.
	//
	// Since v1.5.0.
	AccessToken string `json:"accessToken,omitempty"`
}
