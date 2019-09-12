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

type AccessToken struct {
	//
	AccessTokenHash string `json:"accessTokenHash,omitempty"`

	//
	RefresthTokenHash string `json:"refreshTokenHash,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	Subject string `json:"subject,omitempty"`

	//
	GrantType types.GrantType `json:"grantType,omitempty"`

	//
	Scopes []string `json:"scopes,omitempty"`

	//
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt,omitempty"`

	//
	RefreshTokenExpiresAt uint64 `json:"refreshTokenExpiresAt,omitempty"`

	//
	CreatedAt uint64 `json:"createdAt,omitempty"`

	//
	LastRefreshedAt uint64 `json:"lastRefreshedAt,omitempty"`

	//
	Properties []Property `json:"properties,omitempty"`
}
