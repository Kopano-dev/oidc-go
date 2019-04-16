/*
 * Copyright 2019 Kopano
 *
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE.txt file.
 *
 */

package oidc

type logger interface {
	Printf(string, ...interface{})
}

type noopLogger struct {
}

func (log *noopLogger) Printf(string, ...interface{}) {
}

// DefaultLogger is the logger used by this library if no other is explicitly
// specified.
var DefaultLogger logger = &noopLogger{}
