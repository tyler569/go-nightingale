// NGTODO: generate this

// +build amd64,nightingale

package syscall

const (
	// syscall numbers
	// These must stay sorted relative to
	// include/ng/syscall_consts.h
	NG_INVALID = iota
	NG_DEBUGPRINT
	NG_EXIT
	NG_OPEN
	NG_READ
	NG_WRITE
	NG_FORK
	NG_TOP
	NG_GETPID
	NG_GETTID
	NG_EXECVE
	NG_WAIT4
	NG_SOCKET
	NG_BIND0
	NG_CONNECT0
	NG_STRACE
	NG_BIND
	NG_CONNECT
	NG_SEND
	NG_SENDTO
	NG_RECV
	NG_RECVFROM
	NG_WAITPID
	NG_DUP2
	NG_UNAME
	NG_YIELD
	NG_SEEK
	NG_POLL
	NG_MMAP
	NG_MUNMAP
	NG_HEAPDBG
	NG_SETPGID
	NG_EXIT_GROUP
	NG_CLONE0
	NG_LOADMOD
	NG_HALTVM
	NG_OPENAT
	NG_EXECVEAT
	NG_TTYCTL
	NG_CLOSE
	NG_PIPE
	NG_SIGACTION
	NG_SIGRETURN
	NG_KILL
	NG_SLEEPMS
	NG_GETDIRENTS
	NG_TIME
	NG_CREATE
	NG_PROCSTATE
	NG_FAULT
)

// Errors
const (
	SUCCESS = Errno(iota)
	EINVAL
	EAGAIN
	ENOEXEC
	ENOENT
	EAFNOSUPPORT
	EPROTONOSUPPORT
	ECHILD
	EPERM
	EFAULT
	EBADF
	ERANGE
	EDOM
	EACCES
	ESPIPE
	EISDIR
	ENOMEM
	ETODO

	///////

	EWOULDBLOCK = EAGAIN
)

// Signals
const (
	SIGSEGV   = Signal(1)
	SIGINT    = Signal(2)
)

// Error table
var errors = [...]string{
	0: "success",
	1: "invalid argument",
	2: "would block",
	3: "not executable",
	4: "entity does not exist",
	5: "no protocol support",
	6: "no protocol support",
	7: "no such child",
	8: "permission error",
	9: "fault",
	10: "bad file descriptor",
	11: "out of range",
	12: "domain error",
	13: "illegal access",
	14: "operation invalid for pipe",
	15: "operation invalid for directory",
	16: "operation is TODO in kernel",
}

// Signal table
var signals = [...]string{
	1: "segmentation fault",
	2: "interrupt",
}
