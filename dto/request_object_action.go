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

type RequestObjectAction string

const (
	RequestObjectAction_CREATED               = RequestObjectAction(`CREATED`)
	RequestObjectAction_BAD_REQUEST           = RequestObjectAction(`BAD_REQUEST`)
	RequestObjectAction_UNAUTHORIZED          = RequestObjectAction(`UNAUTHORIZED`)
	RequestObjectAction_FORBIDDEN             = RequestObjectAction(`FORBIDDEN`)
	RequestObjectAction_PAYLOAD_TOO_LARGE     = RequestObjectAction(`PAYLOAD_TOO_LARGE`)
	RequestObjectAction_INTERNAL_SERVER_ERROR = RequestObjectAction(`INTERNAL_SERVER_ERROR`)
)
