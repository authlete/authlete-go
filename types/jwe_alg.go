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

type JWEAlg string

const (
	JWEAlg_RSA1_5             = JWEAlg(`RSA1_5`)
	JWEAlg_RSA_OAEP           = JWEAlg(`RSA_OAEP`)
	JWEAlg_RSA_OAEP_256       = JWEAlg(`RSA_OAEP_256`)
	JWEAlg_A128KW             = JWEAlg(`A128KW`)
	JWEAlg_A192KW             = JWEAlg(`A192KW`)
	JWEAlg_A256KW             = JWEAlg(`A256KW`)
	JWEAlg_DIR                = JWEAlg(`DIR`)
	JWEAlg_ECDH_ES            = JWEAlg(`ECDH_ES`)
	JWEAlg_ECDH_ES_A128KW     = JWEAlg(`ECDH_ES_A128KW`)
	JWEAlg_ECDH_ES_A192KW     = JWEAlg(`ECDH_ES_A192KW`)
	JWEAlg_ECDH_ES_A256KW     = JWEAlg(`ECDH_ES_A256KW`)
	JWEAlg_A128GCMKW          = JWEAlg(`A128GCMKW`)
	JWEAlg_A192GCMKW          = JWEAlg(`A192GCMKW`)
	JWEAlg_A256GCMKW          = JWEAlg(`A256GCMKW`)
	JWEAlg_PBES2_HS256_A128KW = JWEAlg(`PBES2_HS256_A128KW`)
	JWEAlg_PBES2_HS384_A192KW = JWEAlg(`PBES2_HS384_A192KW`)
	JWEAlg_PBES2_HS512_A256KW = JWEAlg(`PBES2_HS512_A256KW`)
)
