//
// Copyright (C) 2019-2022 Authlete, Inc.
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

type TokenAction string

const (
	TokenAction_INVALID_CLIENT        = TokenAction(`INVALID_CLIENT`)
	TokenAction_INTERNAL_SERVER_ERROR = TokenAction(`INTERNAL_SERVER_ERROR`)
	TokenAction_BAD_REQUEST           = TokenAction(`BAD_REQUEST`)
	TokenAction_PASSWORD              = TokenAction(`PASSWORD`)
	TokenAction_OK                    = TokenAction(`OK`)
	TokenAction_TOKEN_EXCHANGE        = TokenAction(`TOKEN_EXCHANGE`)
	TokenAction_JWT_BEARER            = TokenAction(`JWT_BEARER`)
)
