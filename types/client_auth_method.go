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

type ClientAuthMethod string

const (
	ClientAuthMethod_NONE                        = ClientAuthMethod(`NONE`)
	ClientAuthMethod_CLIENT_SECRET_BASIC         = ClientAuthMethod(`CLIENT_SECRET_BASIC`)
	ClientAuthMethod_CLIENT_SECRET_POST          = ClientAuthMethod(`CLIENT_SECRET_POST`)
	ClientAuthMethod_CLIENT_SECRET_JWT           = ClientAuthMethod(`CLIENT_SECRET_JWT`)
	ClientAuthMethod_PRIVATE_KEY_JWT             = ClientAuthMethod(`PRIVATE_KEY_JWT`)
	ClientAuthMethod_TLS_CLIENT_AUTH             = ClientAuthMethod(`TLS_CLIENT_AUTH`)
	ClientAuthMethod_SELF_SIGNED_TLS_CLIENT_AUTH = ClientAuthMethod(`SELF_SIGNED_TLS_CLIENT_AUTH`)
)
