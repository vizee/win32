// generated by genapi.go
// GOFILE=ws2_32.go GOPACKAGE=ws2_32
// DO NOT EDIT!
package ws2_32

import (
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer // keep unsafe

var ()

func mustload(libname string) syscall.Handle {
	hlib, err := syscall.LoadLibrary(libname)
	if err != nil {
		panic(err)
	}
	return hlib
}

var (
	pfngetprocaddress uintptr
)

func mustfind(hmodule syscall.Handle, procname string) uintptr {
	ptr := uintptr(0)
	if procname[0] == '#' {
		for i := 1; i < len(procname); i++ {
			c := procname[i]
			if c < '0' || c > '9' {
				break
			}
			ptr = ptr*10 + uintptr(c-'0')
		}
	} else {
		ptr = *(*uintptr)(unsafe.Pointer(&procname))
	}
	proc, _, err := syscall.Syscall(pfngetprocaddress, 2,
		uintptr(hmodule),
		ptr,
		0)
	if proc == 0 {
		panic(err)
	}
	return proc
}

func boolcast(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

func init() {
	hkernel32 := mustload("kernel32.dll")
	var err error
	pfngetprocaddress, err = syscall.GetProcAddress(hkernel32, "GetProcAddress")
	if err != nil {
		panic(err)
	}
	hws2_32 := mustload("ws2_32.dll")
	_ = hws2_32
}
