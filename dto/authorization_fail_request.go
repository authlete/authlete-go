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

// Request to Authlete's /api/auth/authorization/fail API.
type AuthorizationFailRequest struct {
	// The ticket issued by Authlete's /api/auth/authorization API.
	Ticket string `json:"ticket,omitempty"`

	// The reason of the failure of the authorization request.
	Reason AuthorizationFailReason `json:"reason,omitempty"`

	// The custom description about the authorization failure.
	Description string `json:"description,omitempty"`
}
