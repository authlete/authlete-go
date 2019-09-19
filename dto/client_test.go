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
		"developer": "Developer",
		"clientId": 4326385671,
		"clientIdAlias": "MyAlias",
		"clientIdAliasEnabled": true,
		"clientSecret": "Secret",
		"clientType": "PUBLIC",
		"redirectUris": ["https://example.com/redirect0", "https://example.com/redirect1"],
		"responseTypes": ["CODE","TOKEN"],
		"grantTypes": ["AUTHORIZATION_CODE","IMPLICIT"],
		"applicationType": "WEB",
		"contacts": ["info@example.com", "admin@example.com"],
		"clientName": "Apple",
		"clientNames": [
			{"tag":"ja", "value":"Ringo"},
			{"tag":"de", "value":"Apfel"}
		]
	}`

	var client Client

	err := json.Unmarshal([]byte(jsn), &client)
	if err != nil {
		t.Error(`Cannot unmarshal the input JSON.`)
		return
	}

	redirectUris := []string{"https://example.com/redirect0", "https://example.com/redirect1"}
	responseTypes := []types.ResponseType{types.ResponseType_CODE, types.ResponseType_TOKEN}
	grantTypes := []types.GrantType{types.GrantType_AUTHORIZATION_CODE, types.GrantType_IMPLICIT}
	contacts := []string{"info@example.com", "admin@example.com"}
	clientNames := []TaggedValue{TaggedValue{`ja`, `Ringo`}, TaggedValue{`de`, `Apfel`}}

	assert.Equal(t, `Developer`, client.Developer)
	assert.Equal(t, uint64(4326385671), client.ClientId)
	assert.Equal(t, `MyAlias`, client.ClientIdAlias)
	assert.Equal(t, true, client.ClientIdAliasEnabled)
	assert.Equal(t, `Secret`, client.ClientSecret)
	assert.Equal(t, types.ClientType_PUBLIC, client.ClientType)
	assert.Equal(t, redirectUris, client.RedirectUris)
	assert.Equal(t, responseTypes, client.ResponseTypes)
	assert.Equal(t, grantTypes, client.GrantTypes)
	assert.Equal(t, types.ApplicationType_WEB, client.ApplicationType)
	assert.Equal(t, contacts, client.Contacts)
	assert.Equal(t, `Apple`, client.ClientName)
	assert.Equal(t, clientNames, client.ClientNames)
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
