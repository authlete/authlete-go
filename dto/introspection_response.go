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

type IntrospectionResponse struct {
	ApiResponse

	//
	Action IntrospectionAction `json:"action,omitempty"`

	//
	ClientId uint64 `json:"clientId,omitempty"`

	//
	Subject string `json:"subject,omitempty"`

	//
	Scopes []string `json:"scopes,omitempty"`

	//
	Existent bool `json:"existent,omitempty"`

	//
	Usable bool `json:"usable,omitempty"`

	//
	Sufficient bool `json:"sufficient,omitempty"`

	//
	Refreshable bool `json:"refreshable,omitempty"`

	//
	ResponseContent string `json:"responseContent,omitempty"`

	//
	ExpiresAt uint64 `json:"expiresAt,omitempty"`

	//
	Properties []Property `json:"properties,omitempty"`

	//
	ClientIdAlias string `json:"clientIdAlias,omitempty"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed,omitempty"`

	//
	CertificateThumbprint string `json:"certificateThumbprint,omitempty"`

	//
	Resources []string `json:"resources,omitempty"`

	//
	AccessTokenResources []string `json:"accessTokenResources,omitempty"`

	//
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`

	//
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
}
