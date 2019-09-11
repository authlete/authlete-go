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
	AccessTokenHash string `json:"accessTokenHash"`

	//
	RefresthTokenHash string `json:"refreshTokenHash"`

	//
	ClientId uint64 `json:"clientId"`

	//
	Subject string `json:"subject"`

	//
	GrantType types.GrantType `json:"grantType"`

	//
	Scopes []string `json:"scopes"`

	//
	AccessTokenExpiresAt uint64 `json:"accessTokenExpiresAt"`

	//
	RefreshTokenExpiresAt uint64 `json:"refreshTokenExpiresAt"`

	//
	CreatedAt uint64 `json:"createdAt"`

	//
	LastRefreshedAt uint64 `json:"lastRefreshedAt"`

	//
	Properties []Property `json:"properties"`
}
