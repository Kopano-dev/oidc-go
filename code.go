/*
 * Copyright 2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
)

// Code challenge methods implemented by Konnect. See https://tools.ietf.org/html/rfc7636.
const (
	PlainCodeChallengeMethod = "plain"
	S256CodeChallengeMethod  = "S256"
)

// ValidateCodeChallenge implements https://tools.ietf.org/html/rfc7636#section-4.6
// code challenge verification.
func ValidateCodeChallenge(challenge string, method string, verifier string) error {
	if method == "" {
		// We default to S256CodeChallengeMethod.
		method = S256CodeChallengeMethod
	}

	computed, err := MakeCodeChallenge(method, verifier)
	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare([]byte(challenge), []byte(computed)) != 1 {
		return errors.New("invalid code challenge")
	}
	return nil
}

// MakeCodeChallenge creates a code challenge using the provided method and
// verifier for https://tools.ietf.org/html/rfc7636#section-4.6 verification.
func MakeCodeChallenge(method string, verifier string) (string, error) {
	if verifier == "" {
		return "", errors.New("invalid verifier")
	}

	switch method {
	case PlainCodeChallengeMethod:
		// Challenge is verifier.
		return verifier, nil
	case S256CodeChallengeMethod:
		// BASE64URL-ENCODE(SHA256(ASCII(code_verifier)))
		sum := sha256.Sum256([]byte(verifier))
		return base64.URLEncoding.EncodeToString(sum[:]), nil
	}

	return "", errors.New("transform algorithm not supported")
}
