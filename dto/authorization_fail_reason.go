//
// Copyright (C) 2019-2020 Authlete, Inc.
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

type AuthorizationFailReason string

const (
	// UNKNOWN
	//
	// Unknown reason.
	AuthorizationFailReason_UNKNOWN = AuthorizationFailReason(`UNKNOWN`)

	// NOT_LOGGED_IN
	//
	// The authorization request from the client application contained
	// `prompt=none`, but any end-user has not logged in.
	//
	// Using this reason will result in `error=login_required`.
	AuthorizationFailReason_NOT_LOGGED_IN = AuthorizationFailReason(`NOT_LOGGED_IN`)

	// MAX_AGE_NOT_SUPPORTED
	//
	// The authorization request from the client application contained
	// `max_age` parameter with a non-zero value or the client's
	// configuration has a non-zero value for `default_max_age`
	// configuration parameter, but the authorization server cannot
	// behave properly based on the max age value mainly because the
	// server does not manage authentication time of end-users.
	//
	// Using this reason will result in `error=login_required`.
	AuthorizationFailReason_MAX_AGE_NOT_SUPPORTED = AuthorizationFailReason(`MAX_AGE_NOT_SUPPORTED`)

	// EXCEEDS_MAX_AGE
	//
	// The authorization request from the client application contained
	// `prompt=none`, but the time specified by `max_age` request parameter
	// or by `default_max_age` configuration parameter has passed since the
	// time at which the end-user logged in.
	//
	// Using this reason will result in `error=login_required`.
	AuthorizationFailReason_EXCEEDS_MAX_AGE = AuthorizationFailReason(`EXCEEDS_MAX_AGE`)

	// DIFFERENT_SUBJECT
	//
	// The authorization request from the client application requested a
	// specific value for the `sub` claim, but the current end-user (in
	// the case of `prompt=none`) or the end-user after the authentication
	// is different from the specified value.
	//
	// Using this reason will result in `error=login_required`.
	AuthorizationFailReason_DIFFERENT_SUBJECT = AuthorizationFailReason(`DIFFERENT_SUBJECT`)

	// ACR_NOT_SATISFIED
	//
	// The authorization request from the client application contained the
	// `acr` claim in the `claims` request parameter and the claim was marked
	// as essential, but the ACR performed for the end-user does not match
	// any one of the requested ACRs.
	//
	// Using this reason will result in `error=login_required`.
	AuthorizationFailReason_ACR_NOT_SATISFIED = AuthorizationFailReason(`ACR_NOT_SATISFIED`)

	// DENINED
	//
	// The end-user denied the authorization request from the client application.
	//
	// Using this reason will result in `error=access_denied`.
	AuthorizationFailReason_DENIED = AuthorizationFailReason(`DENIED`)

	// SERVER_ERROR
	//
	// Server error.
	//
	// Using this reason will result in `error=server_error`.
	AuthorizationFailReason_SERVER_ERROR = AuthorizationFailReason(`SERVER_ERROR`)

	// NOT_AUTHENTICATED
	//
	// The end-user was not authenticated.
	//
	// Using this reason will result in `error=login_required`.
	AuthorizationFailReason_NOT_AUTHENTICATED = AuthorizationFailReason(`NOT_AUTHENTICATED`)

	// ACCOUNT_SELECTION_REQUIRED
	//
	// The authorization server cannot obtain an account selection choice
	// made by the end-user.
	//
	// Using this reason will result in `error=account_selection_required`.
	AuthorizationFailReason_ACCOUNT_SELECTION_REQUIRED = AuthorizationFailReason(`ACCOUNT_SELECTION_REQUIRED`)

	// CONSENT_REQUIRED
	//
	// The authorization server cannot obtain consent from the end-user.
	//
	// Using this reason will result in `error=consent_required`.
	AuthorizationFailReason_CONSENT_REQUIRED = AuthorizationFailReason(`CONSENT_REQUIRED`)

	// INTERACTION_REQUIRED
	//
	// The authorization server needs interaction with the end-user.
	//
	// Using this reason will result in `error=interaction_required`.
	AuthorizationFailReason_INTERACTION_REQUIRED = AuthorizationFailReason(`INTERACTION_REQUIRED`)

	// INVALID_TARGET
	//
	// The requested resource is invalid, missing, unknown, or malformed.
	// See RFC 8707 Resource Indicators for OAuth 2.0 for details.
	//
	// Using this reason will result in `error=invalid_target`.
	//
	// Since v1.1.0.
	AuthorizationFailReason_INVALID_TARGET = AuthorizationFailReason(`INVALID_TARGET`)
)
