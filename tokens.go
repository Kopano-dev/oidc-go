/*
 * Copyright 2017-2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

// TokenTypeBearer is required for OIDC as defined in http://openid.net/specs/openid-connect-core-1_0.html.
const TokenTypeBearer = "Bearer"

// Token header as used in JSON web tokens.
const (
	JWTHeaderKeyID = "kid"
	JWTHeaderAlg   = "alg"
)
