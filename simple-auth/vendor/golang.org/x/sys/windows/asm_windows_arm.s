// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// 

#include "textflag.h"

TEXT ·getprocaddress(SB),NOSPLIT,$0
	B	syscall·getprocaddress(SB)

TEXT ·loadlibrary(SB),NOSPLIT,$0
	B	syscall·loadlibrary(SB)
