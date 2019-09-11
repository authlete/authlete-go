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

type ResponseType string

const (
	ResponseType_NONE                = ResponseType(`NONE`)
	ResponseType_CODE                = ResponseType(`CODE`)
	ResponseType_TOKEN               = ResponseType(`TOKEN`)
	ResponseType_ID_TOKEN            = ResponseType(`ID_TOKEN`)
	ResponseType_CODE_TOKEN          = ResponseType(`CODE_TOKEN`)
	ResponseType_CODE_ID_TOKEN       = ResponseType(`CODE_ID_TOKEN`)
	ResponseType_ID_TOKEN_TOKEN      = ResponseType(`ID_TOKEN_TOKEN`)
	ResponseType_CODE_ID_TOKEN_TOKEN = ResponseType(`CODE_ID_TOKEN_TOKEN`)
)
