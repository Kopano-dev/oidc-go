/*
 * Copyright 2019 Kopano
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
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
