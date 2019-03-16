// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// 

package terminal

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS
const ioctlWriteTermios = unix.TCSETS
