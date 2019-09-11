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

import "os"

type AuthleteEnvConfiguration struct {
}

func (self *AuthleteEnvConfiguration) GetBaseUrl() string {
	return os.Getenv(`AUTHLETE_BASE_URL`)
}

func (self *AuthleteEnvConfiguration) GetServiceOwnerApiKey() string {
	return os.Getenv(`AUTHLETE_SERVICEOWNER_APIKEY`)
}

func (self *AuthleteEnvConfiguration) GetServiceOwnerApiSecret() string {
	return os.Getenv(`AUTHLETE_SERVICEOWNER_APISECRET`)
}

func (self *AuthleteEnvConfiguration) GetServiceOwnerAccessToken() string {
	return os.Getenv(`AUTHLETE_SERVICEOWNER_ACCESSTOKEN`)
}

func (self *AuthleteEnvConfiguration) GetServiceApiKey() string {
	return os.Getenv(`AUTHLETE_SERVICE_APIKEY`)
}

func (self *AuthleteEnvConfiguration) GetServiceApiSecret() string {
	return os.Getenv(`AUTHLETE_SERVICE_APISECRET`)
}

func (self *AuthleteEnvConfiguration) GetServiceAccessToken() string {
	return os.Getenv(`AUTHLETE_SERVICE_ACCESSTOKEN`)
}
