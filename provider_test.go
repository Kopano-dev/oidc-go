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
	"net/http"
	"testing"
	"time"
)

const baseTimeout = time.Second * 10

var httpTimeoutClient = &http.Client{
	Timeout: baseTimeout,
}

func TestNewProvider(t *testing.T) {
	config := &ProviderConfig{
		Logger:     &testingLogger{t},
		HTTPClient: httpTimeoutClient,
		HTTPHeader: http.Header{},
	}
	config.HTTPHeader.Set("X-Custom-Header", "1")

	ctx := context.Background()

	provider, err := NewProvider(ctx, "https://accounts.google.com", config)
	if err != nil {
		t.Fatal(err)
	}

	updates := make(chan *ProviderDefinition)
	errors := make(chan error)
	err = provider.Initialize(ctx, updates, errors)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-provider.Ready():
		t.Log("ready")
	case <-time.After(baseTimeout):
		t.Error("timeout waiting for ready")
	}

	select {
	case <-updates:
		t.Log("update received")
	case err := <-errors:
		t.Errorf("error received: %v", err)
	case <-time.After(baseTimeout):
		t.Error("timeout waiting for updates")
	}
}
