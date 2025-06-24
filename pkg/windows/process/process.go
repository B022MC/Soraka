package process

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modKernel32                  = syscall.NewLazyDLL("kernel32.dll")
	procCloseHandle              = modKernel32.NewProc("CloseHandle")
	procCreateToolhelp32Snapshot = modKernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32First           = modKernel32.NewProc("Process32FirstW")
	procProcess32Next            = modKernel32.NewProc("Process32NextW")
	queryFullProcessImageName    = modKernel32.NewProc("QueryFullProcessImageNameW")
	errNotFoundProcess           = errors.New("未找到进程")
)

const (
	ERROR_NO_MORE_FILES = 0x12
	MAX_PATH            = 260
)

type PROCESSENTRY32 struct {
	Size              uint32
	CntUsage          uint32
	ProcessID         uint32
	DefaultHeapID     uintptr
	ModuleID          uint32
	CntThreads        uint32
	ParentProcessID   uint32
	PriorityClassBase int32
	Flags             uint32
	ExeFile           [MAX_PATH]uint16
}

type Process struct {
	pid  int
	ppid int
	exe  string
}

func (p *Process) Pid() int           { return p.pid }
func (p *Process) PPid() int          { return p.ppid }
func (p *Process) Executable() string { return p.exe }

func newWindowsProcess(e *PROCESSENTRY32) *Process {
	end := 0
	for e.ExeFile[end] != 0 {
		end++
	}
	return &Process{
		pid:  int(e.ProcessID),
		ppid: int(e.ParentProcessID),
		exe:  syscall.UTF16ToString(e.ExeFile[:end]),
	}
}

func Processes() ([]*Process, error) {
	handle, _, _ := procCreateToolhelp32Snapshot.Call(0x00000002, 0)
	if handle < 0 {
		return nil, syscall.GetLastError()
	}
	defer procCloseHandle.Call(handle)

	var entry PROCESSENTRY32
	entry.Size = uint32(unsafe.Sizeof(entry))
	ret, _, _ := procProcess32First.Call(handle, uintptr(unsafe.Pointer(&entry)))
	if ret == 0 {
		return nil, fmt.Errorf("Error retrieving process info")
	}

	var results []*Process
	for {
		results = append(results, newWindowsProcess(&entry))
		ret, _, _ = procProcess32Next.Call(handle, uintptr(unsafe.Pointer(&entry)))
		if ret == 0 {
			break
		}
	}
	return results, nil
}

func GetProcessFullPath(targetName string) (string, error) {
	pid, err := findPidByName(targetName)
	if err != nil {
		return "", err
	}
	hProcess, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, uint32(pid))
	if err != nil {
		log.Printf("无法获取到进程 handle: %v", err)
		return "", errNotFoundProcess
	}
	defer syscall.CloseHandle(hProcess)

	var buf [MAX_PATH]uint16
	size := uint32(len(buf))
	ret, _, lastErr := queryFullProcessImageName.Call(uintptr(hProcess), 0, uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&size)))
	if ret != 1 {
		errMsg := "none"
		if lastErr != nil {
			errMsg = lastErr.Error()
		}
		return "", errors.New("获取进程全路径失败: " + errMsg)
	}
	return syscall.UTF16ToString(buf[:]), nil
}

func GetProcessCommand(targetName string) (string, error) {
	pid, err := findPidByName(targetName)
	if err != nil {
		return "", err
	}
	return GetCmdline(uint32(pid))
}

func findPidByName(targetName string) (int, error) {
	processList, err := Processes()
	if err != nil {
		return 0, err
	}
	for _, processInfo := range processList {
		if processInfo.Executable() == targetName {
			return processInfo.Pid(), nil
		}
	}
	return 0, errNotFoundProcess
}

func GetCmdline(pid uint32) (string, error) {
	h, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ, false, pid)
	if err != nil {
		if e, ok := err.(windows.Errno); ok && e == windows.ERROR_ACCESS_DENIED {
			return "", nil // 没权限, 忽略
		}
		return "", err
	}
	defer windows.CloseHandle(h)

	var pbi struct {
		ExitStatus                   uint32
		PebBaseAddress               uintptr
		AffinityMask                 uintptr
		BasePriority                 int32
		UniqueProcessID              uintptr
		InheritedFromUniqueProcessID uintptr
	}
	pbiLen := uint32(unsafe.Sizeof(pbi))
	err = windows.NtQueryInformationProcess(h, windows.ProcessBasicInformation, unsafe.Pointer(&pbi), pbiLen, &pbiLen)
	if err != nil {
		return "", err
	}

	var addr uint64
	d := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&addr)),
		Len:  8, Cap: 8}))
	err = windows.ReadProcessMemory(h, pbi.PebBaseAddress+32, &d[0], uintptr(len(d)), nil)
	if err != nil {
		return "", err
	}

	var commandLine windows.NTUnicodeString
	d = *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&commandLine)),
		Len:  int(unsafe.Sizeof(commandLine)), Cap: int(unsafe.Sizeof(commandLine))}))
	err = windows.ReadProcessMemory(h, uintptr(addr+112), &d[0], uintptr(len(d)), nil)
	if err != nil {
		return "", err
	}

	cmdData := make([]uint16, commandLine.Length/2)
	d = *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&cmdData[0])),
		Len:  int(commandLine.Length), Cap: int(commandLine.Length)}))
	err = windows.ReadProcessMemory(h, uintptr(unsafe.Pointer(commandLine.Buffer)), &d[0], uintptr(commandLine.Length), nil)
	if err != nil {
		return "", err
	}

	return windows.UTF16ToString(cmdData), nil
}
