/*
 * Copyright 2017 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

// OIDC response types and flows.
const (
	ResponseTypeCode             = "code"                // OIDC code flow
	ResponseTypeIDTokenToken     = "id_token token"      // OIDC implicit flow
	ResponseTypeIDToken          = "id_token"            // OIDC implicit flow
	ResponseTypeCodeIDToken      = "code id_token"       // OIDC hybrid flow
	ResponseTypeCodeToken        = "code token"          // OIDC hybrid flow
	ResponseTypeCodeIDTokenToken = "code id_token token" // OIDC hybrid flow
	ResponseTypeToken            = "token"               // OAuth2

	ResponseModeFragment = "fragment"
	ResponseModeQuery    = "query"

	FlowCode     = "code"
	FlowImplicit = "implicit"
	FlowHybrid   = "hybrid"
)
