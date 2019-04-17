/*
 * Copyright 2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pquerna/cachecontrol"
)

// Basic HTTP related global settings.
var (
	DefaultHTTPClient       *http.Client
	DefaultMaxJSONFetchSize int64 = 5 * 1024 * 1024 // 5 MiB
	DefaultJSONFetchExpiry        = time.Minute * 1
	DefaultJSONFetchRetry         = time.Second * 3
)

func fetchJSON(ctx context.Context, u *url.URL, dst interface{}, client *http.Client) (time.Duration, error) {
	if client == nil {
		client = DefaultHTTPClient
		if client == nil {
			client = http.DefaultClient
		}
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return DefaultJSONFetchRetry, fmt.Errorf("failed create fetch JSON request: %v", err)
	}
	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return DefaultJSONFetchRetry, fmt.Errorf("failed to fetch JSON: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return DefaultJSONFetchRetry, fmt.Errorf("failed to fetch JSON (status: %d)", res.StatusCode)
	}

	_, expires, _ := cachecontrol.CachableResponse(req, res, cachecontrol.Options{})
	err = json.NewDecoder(io.LimitReader(res.Body, DefaultMaxJSONFetchSize)).Decode(dst)
	if err != nil {
		return DefaultJSONFetchRetry, fmt.Errorf("failed to fetch JSON: %v", err)
	}

	expirationDuration := expires.Sub(time.Now())
	if expirationDuration < DefaultJSONFetchRetry {
		if err == nil {
			expirationDuration = DefaultJSONFetchExpiry
		} else {
			expirationDuration = DefaultJSONFetchRetry
		}
	}

	return expirationDuration, err
}
