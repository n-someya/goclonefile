package goclonefile

import (
	"fmt"
	"runtime"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	OS_DARWIN = "darwin"
)

//Note: go/src/go/build/syslist.go
//package build

//const goosList = "aix android darwin dragonfly freebsd hurd js linux nacl netbsd openbsd plan9 solaris windows zos "
//const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc riscv riscv64 s390 s390x sparc sparc64 wasm "

type CLONEFILE_FLAG uint32

var (
	// Flags
	CLONE_NOFOLLOW    CLONEFILE_FLAG = 0x0001 /* Don't follow symbolic links */
	CLONE_NOOWNERCOPY CLONEFILE_FLAG = 0x0002 /* Don't copy ownership information from source */
	AT_FDCWD                         = -2
)

// Do the interface allocations only once for common
// Errno values.
var (
	errEAGAIN error = unix.EAGAIN
	errEINVAL error = unix.EINVAL
	errENOENT error = unix.ENOENT
)

func errnoErr(e unix.Errno) error {
	switch e {
	case 0:
		return nil
	case unix.EAGAIN:
		return errEAGAIN
	case unix.EINVAL:
		return errEINVAL
	case unix.ENOENT:
		return errENOENT
	}
	return e
}

// Clonefile clonefile is fast copy method for darwin(MacOS)
func Clonefile(src, dst string) (err error) {
	if runtime.GOOS != OS_DARWIN {
		return fmt.Errorf("Clonefile is implemented for macOS")
	}
	var _p0, _p1 *byte
	_p0, err = unix.BytePtrFromString(src)
	if err != nil {
		return
	}
	_p1, err = unix.BytePtrFromString(dst)
	if err != nil {
		return
	}
	_, _, e1 := unix.Syscall6(unix.SYS_CLONEFILEAT, uintptr(AT_FDCWD), uintptr(unsafe.Pointer(_p0)), uintptr(AT_FDCWD), uintptr(unsafe.Pointer(_p1)), uintptr(CLONE_NOFOLLOW), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
