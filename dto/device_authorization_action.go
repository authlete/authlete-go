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

type DeviceAuthorizationAction string

const (
	DeviceAuthorizationAction_OK                    = DeviceAuthorizationAction(`OK`)
	DeviceAuthorizationAction_BAD_REQUEST           = DeviceAuthorizationAction(`BAD_REQUEST`)
	DeviceAuthorizationAction_UNAUTHORIZED          = DeviceAuthorizationAction(`UNAUTHORIZED`)
	DeviceAuthorizationAction_INTERNAL_SERVER_ERROR = DeviceAuthorizationAction(`INTERNAL_SERVER_ERROR`)
)
