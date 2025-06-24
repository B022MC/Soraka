package utils

import (
	syscall "golang.org/x/sys/windows"
	"os"
	"strings"
	"unsafe"
)

// RelaunchAsAdmin 尝试以管理员身份重新启动当前程序
func RelaunchAsAdmin() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	args := strings.Join(os.Args[1:], " ")

	verb, _ := syscall.UTF16PtrFromString("runas")
	file, _ := syscall.UTF16PtrFromString(exe)
	param, _ := syscall.UTF16PtrFromString(args)
	dir, _ := syscall.UTF16PtrFromString("")
	showCmd := int32(1) // SW_SHOWNORMAL

	r, _, err := syscall.NewLazyDLL("shell32.dll").NewProc("ShellExecuteW").Call(
		0,
		uintptr(unsafe.Pointer(verb)),
		uintptr(unsafe.Pointer(file)),
		uintptr(unsafe.Pointer(param)),
		uintptr(unsafe.Pointer(dir)),
		uintptr(showCmd),
	)
	if r <= 32 {
		return err
	}
	return nil
}

// IsRunAsAdmin 检测当前是否管理员权限
func IsRunAsAdmin() bool {
	var sid *syscall.SID
	// Administrators 组的 SID
	sid, err := syscall.CreateWellKnownSid(syscall.WinBuiltinAdministratorsSid)
	if err != nil {
		return false
	}
	token := syscall.Token(0)
	member, err := token.IsMember(sid)
	if err != nil {
		return false
	}
	return member
}
