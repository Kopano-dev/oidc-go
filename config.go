/*
 * Copyright 2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

import (
	"net/http"
	"net/url"
	"time"
)

// ProviderConfig bundles configuration for a Provider.
type ProviderConfig struct {
	HTTPClient   *http.Client
	Now          func() time.Time
	WellKnownURI *url.URL
	Logger       logger
}

// DefaultProviderConfig is the Provider configuration uses when none was
// explicitly specified.
var DefaultProviderConfig = &ProviderConfig{}
