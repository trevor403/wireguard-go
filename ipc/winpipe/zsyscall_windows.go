// Code generated by 'go generate'; DO NOT EDIT.

package winpipe

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modntdll    = windows.NewLazySystemDLL("ntdll.dll")
	modws2_32   = windows.NewLazySystemDLL("ws2_32.dll")

	procCancelIoEx                         = modkernel32.NewProc("CancelIoEx")
	procConnectNamedPipe                   = modkernel32.NewProc("ConnectNamedPipe")
	procCreateFileW                        = modkernel32.NewProc("CreateFileW")
	procCreateIoCompletionPort             = modkernel32.NewProc("CreateIoCompletionPort")
	procCreateNamedPipeW                   = modkernel32.NewProc("CreateNamedPipeW")
	procGetNamedPipeHandleStateW           = modkernel32.NewProc("GetNamedPipeHandleStateW")
	procGetNamedPipeInfo                   = modkernel32.NewProc("GetNamedPipeInfo")
	procGetQueuedCompletionStatus          = modkernel32.NewProc("GetQueuedCompletionStatus")
	procLocalAlloc                         = modkernel32.NewProc("LocalAlloc")
	procSetFileCompletionNotificationModes = modkernel32.NewProc("SetFileCompletionNotificationModes")
	procNtCreateNamedPipeFile              = modntdll.NewProc("NtCreateNamedPipeFile")
	procRtlDefaultNpAcl                    = modntdll.NewProc("RtlDefaultNpAcl")
	procRtlDosPathNameToNtPathName_U       = modntdll.NewProc("RtlDosPathNameToNtPathName_U")
	procWSAGetOverlappedResult             = modws2_32.NewProc("WSAGetOverlappedResult")
)

func cancelIoEx(file windows.Handle, o *windows.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procCancelIoEx.Addr(), 2, uintptr(file), uintptr(unsafe.Pointer(o)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func connectNamedPipe(pipe windows.Handle, o *windows.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procConnectNamedPipe.Addr(), 2, uintptr(pipe), uintptr(unsafe.Pointer(o)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func createFile(name string, access uint32, mode uint32, sa *windows.SecurityAttributes, createmode uint32, attrs uint32, templatefile windows.Handle) (handle windows.Handle, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	return _createFile(_p0, access, mode, sa, createmode, attrs, templatefile)
}

func _createFile(name *uint16, access uint32, mode uint32, sa *windows.SecurityAttributes, createmode uint32, attrs uint32, templatefile windows.Handle) (handle windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateFileW.Addr(), 7, uintptr(unsafe.Pointer(name)), uintptr(access), uintptr(mode), uintptr(unsafe.Pointer(sa)), uintptr(createmode), uintptr(attrs), uintptr(templatefile), 0, 0)
	handle = windows.Handle(r0)
	if handle == windows.InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func createIoCompletionPort(file windows.Handle, port windows.Handle, key uintptr, threadCount uint32) (newport windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateIoCompletionPort.Addr(), 4, uintptr(file), uintptr(port), uintptr(key), uintptr(threadCount), 0, 0)
	newport = windows.Handle(r0)
	if newport == 0 {
		err = errnoErr(e1)
	}
	return
}

func createNamedPipe(name string, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *windows.SecurityAttributes) (handle windows.Handle, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	return _createNamedPipe(_p0, flags, pipeMode, maxInstances, outSize, inSize, defaultTimeout, sa)
}

func _createNamedPipe(name *uint16, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *windows.SecurityAttributes) (handle windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateNamedPipeW.Addr(), 8, uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(pipeMode), uintptr(maxInstances), uintptr(outSize), uintptr(inSize), uintptr(defaultTimeout), uintptr(unsafe.Pointer(sa)), 0)
	handle = windows.Handle(r0)
	if handle == windows.InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func getNamedPipeHandleState(pipe windows.Handle, state *uint32, curInstances *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32, userName *uint16, maxUserNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetNamedPipeHandleStateW.Addr(), 7, uintptr(pipe), uintptr(unsafe.Pointer(state)), uintptr(unsafe.Pointer(curInstances)), uintptr(unsafe.Pointer(maxCollectionCount)), uintptr(unsafe.Pointer(collectDataTimeout)), uintptr(unsafe.Pointer(userName)), uintptr(maxUserNameSize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getNamedPipeInfo(pipe windows.Handle, flags *uint32, outSize *uint32, inSize *uint32, maxInstances *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetNamedPipeInfo.Addr(), 5, uintptr(pipe), uintptr(unsafe.Pointer(flags)), uintptr(unsafe.Pointer(outSize)), uintptr(unsafe.Pointer(inSize)), uintptr(unsafe.Pointer(maxInstances)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getQueuedCompletionStatus(port windows.Handle, bytes *uint32, key *uintptr, o **ioOperation, timeout uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetQueuedCompletionStatus.Addr(), 5, uintptr(port), uintptr(unsafe.Pointer(bytes)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(o)), uintptr(timeout), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func localAlloc(uFlags uint32, length uint32) (ptr uintptr) {
	r0, _, _ := syscall.Syscall(procLocalAlloc.Addr(), 2, uintptr(uFlags), uintptr(length), 0)
	ptr = uintptr(r0)
	return
}

func setFileCompletionNotificationModes(h windows.Handle, flags uint8) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileCompletionNotificationModes.Addr(), 2, uintptr(h), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ntCreateNamedPipeFile(pipe *windows.Handle, access uint32, oa *objectAttributes, iosb *ioStatusBlock, share uint32, disposition uint32, options uint32, typ uint32, readMode uint32, completionMode uint32, maxInstances uint32, inboundQuota uint32, outputQuota uint32, timeout *int64) (ntstatus error) {
	r0, _, _ := syscall.Syscall15(procNtCreateNamedPipeFile.Addr(), 14, uintptr(unsafe.Pointer(pipe)), uintptr(access), uintptr(unsafe.Pointer(oa)), uintptr(unsafe.Pointer(iosb)), uintptr(share), uintptr(disposition), uintptr(options), uintptr(typ), uintptr(readMode), uintptr(completionMode), uintptr(maxInstances), uintptr(inboundQuota), uintptr(outputQuota), uintptr(unsafe.Pointer(timeout)), 0)
	if r0 != 0 {
		ntstatus = windows.NTStatus(r0)
	}
	return
}

func rtlDefaultNpAcl(dacl *uintptr) (ntstatus error) {
	r0, _, _ := syscall.Syscall(procRtlDefaultNpAcl.Addr(), 1, uintptr(unsafe.Pointer(dacl)), 0, 0)
	if r0 != 0 {
		ntstatus = windows.NTStatus(r0)
	}
	return
}

func rtlDosPathNameToNtPathName(name *uint16, ntName *unicodeString, filePart uintptr, reserved uintptr) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procRtlDosPathNameToNtPathName_U.Addr(), 4, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(ntName)), uintptr(filePart), uintptr(reserved), 0, 0)
	if r0 != 0 {
		ntstatus = windows.NTStatus(r0)
	}
	return
}

func wsaGetOverlappedResult(h windows.Handle, o *windows.Overlapped, bytes *uint32, wait bool, flags *uint32) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procWSAGetOverlappedResult.Addr(), 5, uintptr(h), uintptr(unsafe.Pointer(o)), uintptr(unsafe.Pointer(bytes)), uintptr(_p0), uintptr(unsafe.Pointer(flags)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
