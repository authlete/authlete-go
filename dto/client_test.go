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

	"github.com/authlete/authlete-go/types"
	"github.com/stretchr/testify/assert"
)

func TestClient_FromJson(t *testing.T) {
	jsn := `{
		"clientType": "PUBLIC",
		"applicationType": "WEB"
	}`

	var client Client

	err := json.Unmarshal([]byte(jsn), &client)
	if err != nil {
		t.Error(`Cannot unmarshal the input JSON.`)
		return
	}

	assert.Equal(t, types.ClientType_PUBLIC, client.ClientType)
	assert.Equal(t, types.ApplicationType_WEB, client.ApplicationType)
}

func TestClient_ToJson(t *testing.T) {
	var client Client
	client.ClientType = types.ClientType_PUBLIC
	client.ApplicationType = types.ApplicationType_WEB

	bytes, err := json.Marshal(client)
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

	value, ok := m[`clientType`]
	assert.True(t, ok)
	assert.Equal(t, `PUBLIC`, value)

	value, ok = m[`applicationType`]
	assert.True(t, ok)
	assert.Equal(t, `WEB`, value)
}
