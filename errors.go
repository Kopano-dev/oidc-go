/*
 * Copyright 2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

import (
	"errors"
	"fmt"
)

// A ProviderError is returned for OIDC Provider errors.
type ProviderError struct {
	Err error // The actual error
}

func wrapAsProviderError(err error) error {
	if err == nil {
		return nil
	}

	return &ProviderError{
		Err: err,
	}
}

func (e *ProviderError) Error() string {
	return fmt.Sprintf("oidc provider error: %v", e.Err)
}

// These are the errors that can be returned in ProviderError.Err.
var (
	ErrAllreadyInitialized = errors.New("already initialized")
	ErrNotInitialized      = errors.New("not initialized")
	ErrWrongInitialization = errors.New("wrong initialization")
	ErrIssuerMismatch      = errors.New("issuer mismatch")
)
