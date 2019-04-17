/*
 * Copyright 2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

// OAuth2 error codes.
const (
	ErrorCodeOAuth2UnsupportedResponseType = "unsupported_response_type"
	ErrorCodeOAuth2InvalidRequest          = "invalid_request"
	ErrorCodeOAuth2InvalidToken            = "invalid_token"
	ErrorCodeOAuth2InsufficientScope       = "insufficient_scope"
	ErrorCodeOAuth2InvalidGrant            = "invalid_grant"
	ErrorCodeOAuth2UnsupportedGrantType    = "unsupported_grant_type"
	ErrorCodeOAuth2AccessDenied            = "access_denied"
	ErrorCodeOAuth2ServerError             = "server_error"
	ErrorCodeOAuth2TemporarilyUnavailable  = "temporarily_unavailable"
)

// OIDC error codes.
const (
	ErrorCodeOIDCInteractionRequired = "interaction_required"
	ErrorCodeOIDCLoginRequired       = "login_required"
	ErrorCodeOIDCConsentRequired     = "consent_required"

	ErrorCodeOIDCRequestNotSupported      = "request_not_supported"
	ErrorCodeOIDCInvalidRequestObject     = "invalid_request_object"
	ErrorCodeOIDCRequestURINotSupported   = "request_uri_not_supported"
	ErrorCodeOIDCRegistrationNotSupported = "registration_not_supported"

	ErrorCodeOIDCInvalidRedirectURI    = "invalid_redirect_uri"
	ErrorCodeOIDCInvalidClientMetadata = "invalid_client_metadata"
)
