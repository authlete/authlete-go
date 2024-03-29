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

type TokenUpdateRequest struct {
	//
	AccessToken string `json:"accessToken,omitempty"`

	//
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt,omitempty"`

	//
	Scopes []string `json:"scopes,omitempty"`

	//
	Properties []Property `json:"properties,omitempty"`

	//
	AccessTokenExpiresAtUpdatedOnScopeUpdate bool `json:"accessTokenExpiresAtUpdatedOnScopeUpdate,omitempty"`

	//
	AccessTokenPersistent bool `json:"accessTokenPersistent,omitempty"`

	//
	AccessTokenHash string `json:"accessTokenHash,omitempty"`

	//
	AccessTokenValueUpdated bool `json:"accessTokenValueUpdated,omitempty"`

	//
	CertificateThumbprint string `json:"certificateThumbprint,omitempty"`

	//
	DpopKeyThumbprint string `json:"dpopKeyThumbprint,omitempty"`

	// Flag indicating whether the access token is for an external attachment.
	// See OpenID Connect for Identity Assurance for details.
	//
	// Since v1.1.5.
	ForExternalAttachment bool `json:"forExternalAttachment,omitempty"`

	// Token ID.
	//
	// Since v1.1.5.
	TokenId string `json:"tokenId,omitempty"`
}
