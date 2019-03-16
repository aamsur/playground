// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// 

// +build darwin dragonfly freebsd netbsd openbsd

package terminal

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA
const ioctlWriteTermios = unix.TIOCSETA
