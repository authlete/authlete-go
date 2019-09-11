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

package types

type GrantType string

const (
	GrantType_AUTHORIZATION_CODE = GrantType(`AUTHORIZATION_CODE`)
	GrantType_IMPLICIT           = GrantType(`IMPLICIT`)
	GrantType_PASSWORD           = GrantType(`PASSWORD`)
	GrantType_CLIENT_CREDENTIALS = GrantType(`CLIENT_CREDENTIALS`)
	GrantType_REFRESH_TOKEN      = GrantType(`REFRESH_TOKEN`)
	GrantType_CIBA               = GrantType(`CIBA`)
	GrantType_DEVICE_CODE        = GrantType(`DEVICE_CODE`)
)
