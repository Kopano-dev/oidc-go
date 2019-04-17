/*
 * Copyright 2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

import (
	"testing"
)

type testingLogger struct {
	T *testing.T
}

func (logger *testingLogger) Printf(format string, args ...interface{}) {
	logger.T.Logf(format, args...)
}
