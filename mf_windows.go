//+ build windows

package mediafoundation

import "syscall"

type (
	HRESULT int32
)

type MFFlags byte

const (
	MFSTARTUP_NOSOCKET MFFlags = 0x1
	MFSTARTUP_LITE     MFFlags = 0x1
	MFSTARTUP_FULL     MFFlags = 0
)

var (
	pMFStartup,
	pMECreateAttributes,
	pMFEnumDeviceSources,
	pMFShutdown uintptr
)

func MFStartup(Version uint32, flags MFFlags) HRESULT {
	ret, _, _ := syscall.Syscall(pMFStartup, 2, uintptr(Version), uintptr(flags), 0)
	return HRESULT(ret)
}

func MECreateAttributes(ppMFAttributes *uintptr, cInitialSize uint32) HRESULT {
	return 0
}

func MFEnumDeviceSources(pMFAttributes uintptr, pppSourceActivate **uintptr, pcSourceActivate uint32) HRESULT {
	return 0
}

func MFShutdown() HRESULT {
	ret, _, _ := syscall.Syscall(pMFShutdown, 0, 0, 0, 0)
	return HRESULT(ret)
}

func init() {
	mf, err := syscall.LoadLibrary("Mfplat.dll")
	if err != nil {
		panic("LoadLibrary " + err.Error())
	}
	defer syscall.FreeLibrary(mf)

	pMFStartup = getProcAddr(mf, "MFStartup")
	pMECreateAttributes = getProcAddr(mf, "MECreateAttributes")
	pMFEnumDeviceSources = getProcAddr(mf, "MFEnumDeviceSources")
	pMFShutdown = getProcAddr(mf, "MFShutdown")
}

func getProcAddr(lib syscall.Handle, name string) uintptr {
	addr, err := syscall.GetProcAddress(lib, name)
	if err != nil {
		panic(name + " " + err.Error())
	}
	return addr
}
