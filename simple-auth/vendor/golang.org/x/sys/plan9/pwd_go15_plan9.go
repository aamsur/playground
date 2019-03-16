// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// 

// +build go1.5

package plan9

import "syscall"

func fixwd() {
	syscall.Fixwd()
}

func Getwd() (wd string, err error) {
	return syscall.Getwd()
}

func Chdir(path string) error {
	return syscall.Chdir(path)
}
