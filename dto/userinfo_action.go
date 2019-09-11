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

type UserInfoAction string

const (
	UserInfoAction_INTERNAL_SERVER_ERROR = UserInfoAction(`INTERNAL_SERVER_ERROR`)
	UserInfoAction_BAD_REQUEST           = UserInfoAction(`BAD_REQUEST`)
	UserInfoAction_UNAUTHORIZED          = UserInfoAction(`UNAUTHORIZED`)
	UserInfoAction_FORBIDDEN             = UserInfoAction(`FORBIDDEN`)
	UserInfoAction_OK                    = UserInfoAction(`OK`)
)
