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

type UserInfoResponse struct {
	ApiResponse

	//
	Action UserInfoAction `json:"action"`

	//
	ClientId uint64 `json:"clientId"`

	//
	Subject string `json:"subject"`

	//
	Scopes []string `json:"scopes"`

	//
	Claims []string `json:"claims"`

	//
	Token string `json:"token"`

	//
	ResponseContent string `json:"responseContent"`

	//
	Properties []Property `json:"properties"`

	//
	ClientIdAlias string `json:"clientIdAlias"`

	//
	ClientIdAliasUsed bool `json:"clientIdAliasUsed"`
}
