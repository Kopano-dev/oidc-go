/*
 * Copyright 2017-2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

// OIDC prompt values.
// See http://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
const (
	PromptNone          = "none"
	PromptLogin         = "login"
	PromptConsent       = "consent"
	PromptSelectAccount = "select_account"
)
