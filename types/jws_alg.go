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

type JWSAlg string

const (
	JWSAlg_NONE  = JWSAlg(`NONE`)
	JWSAlg_HS256 = JWSAlg(`HS256`)
	JWSAlg_HS384 = JWSAlg(`HS384`)
	JWSAlg_HS512 = JWSAlg(`HS512`)
	JWSAlg_RS256 = JWSAlg(`RS256`)
	JWSAlg_RS384 = JWSAlg(`RS384`)
	JWSAlg_RS512 = JWSAlg(`RS512`)
	JWSAlg_ES256 = JWSAlg(`ES256`)
	JWSAlg_ES384 = JWSAlg(`ES384`)
	JWSAlg_ES512 = JWSAlg(`ES512`)
	JWSAlg_PS256 = JWSAlg(`PS256`)
	JWSAlg_PS384 = JWSAlg(`PS384`)
	JWSAlg_PS512 = JWSAlg(`PS512`)
)
