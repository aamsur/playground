// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// 

// +build ignore

// mkasm_darwin.go generates assembly trampolines to call libSystem routines from Go.
//This program must be run after mksyscall.go.
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	in1, err := ioutil.ReadFile("syscall_darwin.go")
	if err != nil {
		log.Fatalf("can't open syscall_darwin.go: %s", err)
	}
	arch := os.Args[1]
	in2, err := ioutil.ReadFile(fmt.Sprintf("syscall_darwin_%s.go", arch))
	if err != nil {
		log.Fatalf("can't open syscall_darwin_%s.go: %s", arch, err)
	}
	in3, err := ioutil.ReadFile(fmt.Sprintf("zsyscall_darwin_%s.go", arch))
	if err != nil {
		log.Fatalf("can't open zsyscall_darwin_%s.go: %s", arch, err)
	}
	in := string(in1) + string(in2) + string(in3)

	trampolines := map[string]bool{}

	var out bytes.Buffer

	fmt.Fprintf(&out, "// go run mkasm_darwin.go %s\n", strings.Join(os.Args[1:], " "))
	fmt.Fprintf(&out, "// Code generated by the command above; DO NOT EDIT.\n")
	fmt.Fprintf(&out, "\n")
	fmt.Fprintf(&out, "// +build go1.12\n")
	fmt.Fprintf(&out, "\n")
	fmt.Fprintf(&out, "#include \"textflag.h\"\n")
	for _, line := range strings.Split(in, "\n") {
		if !strings.HasPrefix(line, "func ") || !strings.HasSuffix(line, "_trampoline()") {
			continue
		}
		fn := line[5 : len(line)-13]
		if !trampolines[fn] {
			trampolines[fn] = true
			fmt.Fprintf(&out, "TEXT ·%s_trampoline(SB),NOSPLIT,$0-0\n", fn)
			fmt.Fprintf(&out, "\tJMP\t%s(SB)\n", fn)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf("zsyscall_darwin_%s.s", arch), out.Bytes(), 0644)
	if err != nil {
		log.Fatalf("can't write zsyscall_darwin_%s.s: %s", arch, err)
	}
}
