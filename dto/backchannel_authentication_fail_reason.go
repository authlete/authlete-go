//
// Copyright (C) 2019-2020 Authlete, Inc.
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

type BackchannelAuthenticationFailReason string

const (
	BackchannelAuthenticationFailReason_EXPIRED_LOGIN_HINT_TOKEN = BackchannelAuthenticationFailReason(`EXPIRED_LOGIN_HINT_TOKEN`)
	BackchannelAuthenticationFailReason_UNKNOWN_USER_ID          = BackchannelAuthenticationFailReason(`UNKNOWN_USER_ID`)
	BackchannelAuthenticationFailReason_UNAUTHORIZED_CLIENT      = BackchannelAuthenticationFailReason(`UNAUTHORIZED_CLIENT`)
	BackchannelAuthenticationFailReason_MISSING_USER_CODE        = BackchannelAuthenticationFailReason(`MISSING_USER_CODE`)
	BackchannelAuthenticationFailReason_INVALID_USER_CODE        = BackchannelAuthenticationFailReason(`INVALID_USER_CODE`)
	BackchannelAuthenticationFailReason_INVALID_BINDING_MESSAGE  = BackchannelAuthenticationFailReason(`INVALID_BINDING_MESSAGE`)
	BackchannelAuthenticationFailReason_INVALID_TARGET           = BackchannelAuthenticationFailReason(`INVALID_TARGET`)
	BackchannelAuthenticationFailReason_ACCESS_DENIED            = BackchannelAuthenticationFailReason(`ACCESS_DENIED`)
	BackchannelAuthenticationFailReason_SERVER_ERROR             = BackchannelAuthenticationFailReason(`SERVER_ERROR`)
)
