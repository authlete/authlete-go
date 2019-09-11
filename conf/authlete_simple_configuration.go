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

package conf

type AuthleteSimpleConfiguration struct {
	baseUrl                 string
	serviceOwnerApiKey      string
	serviceOwnerApiSecret   string
	serviceOwnerAccessToken string
	serviceApiKey           string
	serviceApiSecret        string
	serviceAccessToken      string
}

func (self *AuthleteSimpleConfiguration) GetBaseUrl() string {
	return self.baseUrl
}

func (self *AuthleteSimpleConfiguration) SetBaseUrl(url string) {
	self.baseUrl = url
}

func (self *AuthleteSimpleConfiguration) GetServiceOwnerApiKey() string {
	return self.serviceOwnerApiKey
}

func (self *AuthleteSimpleConfiguration) SetServiceOwnerApiKey(key string) {
	self.serviceOwnerApiKey = key
}

func (self *AuthleteSimpleConfiguration) GetServiceOwnerApiSecret() string {
	return self.serviceOwnerApiSecret
}

func (self *AuthleteSimpleConfiguration) SetServiceOwnerApiSecret(secret string) {
	self.serviceOwnerApiSecret = secret
}

func (self *AuthleteSimpleConfiguration) GetServiceOwnerAccessToken() string {
	return self.serviceOwnerAccessToken
}

func (self *AuthleteSimpleConfiguration) SetServiceOwnerAccessToken(token string) {
	self.serviceOwnerAccessToken = token
}

func (self *AuthleteSimpleConfiguration) GetServiceApiKey() string {
	return self.serviceApiKey
}

func (self *AuthleteSimpleConfiguration) SetServiceApiKey(key string) {
	self.serviceApiKey = key
}

func (self *AuthleteSimpleConfiguration) GetServiceApiSecret() string {
	return self.serviceApiSecret
}

func (self *AuthleteSimpleConfiguration) SetServiceApiSecret(secret string) {
	self.serviceApiSecret = secret
}

func (self *AuthleteSimpleConfiguration) GetServiceAccessToken() string {
	return self.serviceAccessToken
}

func (self *AuthleteSimpleConfiguration) SetServiceAccessToken(token string) {
	self.serviceAccessToken = token
}
