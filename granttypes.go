/*
 * Copyright 2017-2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

const (
	// GrantTypeAuthorizationCode is the string value for the
	// OAuth2 authroization code token request grant type.
	GrantTypeAuthorizationCode = "authorization_code"

	// GrantTypeImplicit is the string value for the OAuth2 id_token, token
	// id_token token request grant type.
	GrantTypeImplicit = "implicit"

	// GrantTypeRefreshToken is the string value for the OAuth2 refresh_token
	// token request grant_type.
	GrantTypeRefreshToken = "refresh_token"
)
