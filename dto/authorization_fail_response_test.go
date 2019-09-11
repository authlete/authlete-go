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

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizationFailResponse_FromJson(t *testing.T) {
	jsn := `{
		"resultCode":      "CODE",
		"resultMessage":   "MESSAGE",
		"action":          "LOCATION",
		"responseContent": "CONTENT"
	}`

	var res AuthorizationFailResponse

	err := json.Unmarshal([]byte(jsn), &res)
	if err != nil {
		t.Error(`Cannot unmarshal the input JSON.`)
		return
	}

	assert.Equal(t, `CODE`, res.ResultCode)
	assert.Equal(t, `MESSAGE`, res.ResultMessage)
	assert.Equal(t, AuthorizationFailAction_LOCATION, res.Action)
	assert.Equal(t, `CONTENT`, res.ResponseContent)
}

func TestAuthorizationFailResponse_ToJson(t *testing.T) {
	var res AuthorizationFailResponse
	res.ResultCode = `CODE`
	res.ResultMessage = `MESSAGE`
	res.Action = AuthorizationFailAction_LOCATION
	res.ResponseContent = `CONTENT`

	bytes, err := json.Marshal(res)
	if err != nil {
		t.Error(`Cannot marshal the input data.`)
		return
	}

	// Convert to a generic map.
	var f interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		t.Error(`Cannot unmarshal the input data.`)
		return
	}
	m := f.(map[string]interface{})

	value, ok := m[`resultCode`]
	assert.True(t, ok)
	assert.Equal(t, `CODE`, value)

	value, ok = m[`resultMessage`]
	assert.True(t, ok)
	assert.Equal(t, `MESSAGE`, value)

	value, ok = m[`action`]
	assert.True(t, ok)
	assert.Equal(t, `LOCATION`, value)

	value, ok = m[`responseContent`]
	assert.True(t, ok)
	assert.Equal(t, `CONTENT`, value)
}
