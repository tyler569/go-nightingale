// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Nightingale system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and
// wrap it in our own nicer implementation.

package syscall

import (
)

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err ErrorString)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err ErrorString)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2, err uintptr)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr)

