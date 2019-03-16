// Copyright 2013 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// 

// +build !appengine

package cloudsql

import (
	"errors"
	"net"
)

func connect(instance string) (net.Conn, error) {
	return nil, errors.New(`cloudsql: not supported in App Engine "flexible environment"`)
}