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

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

//go:nosplit
func atoi(b []byte) (n uint) {
	n = 0
	for i := 0; i < len(b); i++ {
		n = n*10 + uint(b[i]-'0')
	}
	return
}

func cstring(s []byte) string {
	for i := range s {
		if s[i] == 0 {
			return string(s[0:i])
		}
	}
	return string(s)
}

func Pipe(p []int) (err error) {
	if len(p) != 2 {
		return NewError("bad arg in system call")
	}
	var pp [2]int32
	err = pipe(&pp)
	p[0] = int(pp[0])
	p[1] = int(pp[1])
	return
}

func Getegid() (egid int) { return -1 }
func Geteuid() (euid int) { return -1 }
func Getgid() (gid int)   { return -1 }
func Getuid() (uid int)   { return -1 }

// NGTODO: move to _types ?
type Signal int
type Sighandler func(int)int

//sys	Exit(code int) (err error)
//sys	ExitGroup(code int) (err error)
//sys	Read(fd int, b []byte) (n int, err error)
//sys	Write(fd int, b []byte) (n int, err error)
//sys	Fork() (pid int32)
//sys	Getpid() (pid int32)
//sys	Gettid() (tid int32)
//sys	Execve(program string, argv *string, envp *string) (err error)
//sys	Socket(domain int, type int, protocol int) (fd int, err error)
//sys	strace(enable bool) (err error)
//	Connect(fd int, addr *Sockaddr, len socklen) (err error)
//	bind
//	send
//	sendto
//	recv
//	recvfrom
//sys	Waitpid(pid int32, status *int, options int) (pid int32, err error)
//sys	Dup2(oldfd int, newfd int) (err error)
//sys	Uname(name Utsname) (err error)
//sys	Yield() (err error)
//sys	Open(path string, mode int) (fd int, err error)
//sys	Seek(fd int, offset int64, whence int) (res int64, err error)
// //sys	Poll()
//sys	Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)
//sys	Munmap(r []byte) (err error)
//sys	Heapdbg(type int) (err error)
//sys	Setpgid(pid int, pgid int) (err error)
//sys	Top(threads bool) (err error)
//	Clone(fn *func(unsafe.Pointer) int, arg unsafe.Pointer, new_stack unsafe.Pointer, flags int) (err error) = NG_CLONE0
//	Openat
//	Execveat
//	Ttyctl
//sys	Clone(fd int) (err error)
//sys	pipe(p *[2]int32) (err error)
//sys	Sigaction(sig Signal, handler Sighandler, flags int) (err error)
//sys	Kill(pid int, sig Signal) (err error)
//sys	sleep(ms int) (err error) = SYS_SLEEPMS
//sys	getdirents(fd int, buf []byte) (n int, err error)
func Getdirents(fd int, buf []Dirent) (n int, err error) {
	return nil, Errno(ETODO)
}


