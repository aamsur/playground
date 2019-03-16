// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// 

// +build windows
// +build !go1.3

package svc

// from go12.c
func getServiceMain(r *uintptr)
