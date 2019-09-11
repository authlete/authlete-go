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

type TokenUpdateRequest struct {
	//
	AccessToken string `json:"accessToken"`

	//
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt"`

	//
	Scopes []string `json:"scopes"`

	//
	Properties []Property `json:"properties"`

	//
	AccessTokenExpiresAtUpdatedOnScopeUpdate bool `json:"accessTokenExpiresAtUpdatedOnScopeUpdate"`

	//
	AccessTokenPersistent bool `json:"accessTokenPersistent"`

	//
	AccessTokenHash string `json:"accessTokenHash"`

	//
	AccessTokenValueUpdated bool `json:"accessTokenValueUpdated"`
}
