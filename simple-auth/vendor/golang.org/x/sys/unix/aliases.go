// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// 

// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
// +build go1.9

package unix

import "syscall"

type Signal = syscall.Signal
type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
