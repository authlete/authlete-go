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

type JWEEnc string

const (
	JWEEnc_A128CBC_HS256 = JWEEnc(`A128CBC_HS256`)
	JWEEnc_A192CBC_HS384 = JWEEnc(`A192CBC_HS384`)
	JWEEnc_A256CBC_HS512 = JWEEnc(`A256CBC_HS512`)
	JWEEnc_A128GCM       = JWEEnc(`A128GCM`)
	JWEEnc_A192GCM       = JWEEnc(`A192GCM`)
	JWEEnc_A256GCM       = JWEEnc(`A256GCM`)
)
