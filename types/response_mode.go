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

type ResponseMode string

const (
	ResponseMode_QUERY         = ResponseMode(`QUERY`)
	ResponseMode_FRAGMENT      = ResponseMode(`FRAGMENT`)
	ResponseMode_FORM_POST     = ResponseMode(`FORM_POST`)
	ResponseMode_JWT           = ResponseMode(`JWT`)
	ResponseMode_QUERY_JWT     = ResponseMode(`QUERY_JWT`)
	ResponseMode_FRAGMENT_JWT  = ResponseMode(`FRAGMENT_JWT`)
	ResponseMode_FORM_POST_JWT = ResponseMode(`FORM_POST_JWT`)
)
