/*
 * Copyright 2017-2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

// Standard claims as used in JSON Web Tokens.
const (
	IssuerIdentifierClaim  = "iss"
	SubjectIdentifierClaim = "sub"
	AudienceClaim          = "aud"
	ExpirationClaim        = "exp"
	IssuedAtClaim          = "iat"
)

// Additional claims as defined by OIDC.
const (
	NameClaim              = "name"
	FamilyNameClaim        = "family_name"
	GivenNameClaim         = "given_name"
	MiddleNameClaim        = "middle_name"
	NicknameClaim          = "nickname"
	PreferredUsernameClaim = "preferred_username"
	ProfileClaim           = "profile"
	PictureClaim           = "picture"
	WebsiteClaim           = "website"
	GenderClaim            = "gender"
	BirthdateClaim         = "birthdate"
	ZoneinfoClaim          = "zoneinfo"
	LocaleClaim            = "locale"
	UpdatedAtClaim         = "updated_at"

	EmailClaim         = "email"
	EmailVerifiedClaim = "email_verified"

	AuthTimeClaim = "auth_time"
)

// Additional claims as defined by OIDC extensions.
const (
	SessionIDClaim = "sid"
)
