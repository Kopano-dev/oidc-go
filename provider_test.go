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
	"testing"
)

func TestNewProvider(t *testing.T) {
	config := &ProviderConfig{
		Logger: &testingLogger{t},
	}

	ctx := context.Background()

	provider, err := NewProvider(ctx, "https://accounts.google.com", config)
	if err != nil {
		t.Fatal(err)
	}

	updates := make(chan *ProviderDefinition)
	err = provider.Initialize(ctx, updates, nil)
	if err != nil {
		t.Fatal(err)
	}

	<-provider.WaitUntilReady()
	t.Log("ready")

	<-updates
	t.Log("update received")
}
