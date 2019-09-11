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

func TestAddress_FromJson(t *testing.T) {
	jsn := `{
		"formatted":      "F",
		"street_address": "S",
		"locality":       "L",
		"region":         "R",
		"postal_code":    "P",
		"country":        "C"
	}`

	var address Address

	err := json.Unmarshal([]byte(jsn), &address)
	if err != nil {
		t.Error(`Cannot unmarshal the input JSON.`)
		return
	}

	assert.Equal(t, `F`, address.Formatted)
	assert.Equal(t, `S`, address.StreetAddress)
	assert.Equal(t, `L`, address.Locality)
	assert.Equal(t, `R`, address.Region)
	assert.Equal(t, `P`, address.PostalCode)
	assert.Equal(t, `C`, address.Country)
}

func TestAddress_ToJson(t *testing.T) {
	address := Address{`F`, `S`, `L`, `R`, `P`, `C`}

	bytes, err := json.Marshal(address)
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

	value, ok := m[`formatted`]
	assert.True(t, ok)
	assert.Equal(t, `F`, value)

	value, ok = m[`street_address`]
	assert.True(t, ok)
	assert.Equal(t, `S`, value)

	value, ok = m[`locality`]
	assert.True(t, ok)
	assert.Equal(t, `L`, value)

	value, ok = m[`region`]
	assert.True(t, ok)
	assert.Equal(t, `R`, value)

	value, ok = m[`postal_code`]
	assert.True(t, ok)
	assert.Equal(t, `P`, value)

	value, ok = m[`country`]
	assert.True(t, ok)
	assert.Equal(t, `C`, value)
}
