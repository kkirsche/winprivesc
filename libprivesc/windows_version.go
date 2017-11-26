package libprivesc

import (
	"fmt"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows"
)

type WindowsVersion struct {
	Major   byte
	Minor   uint8
	Build   uint16
	Version string
}

func GetWindowsVersion(o *OutputConfig) WindowsVersion {
	var w WindowsVersion
	dll := windows.NewLazySystemDLL("kernel32.dll")
	if err := dll.Load(); err != nil {
		fmt.Println("[!] Failed to load kernel32.dll to get windows version")
		return w
	}
	p := dll.NewProc("GetVersion")
	if err := p.Find(); /**/ err != nil {
		fmt.Println("[!] Failed to get version from kernel32.dll")
		return w
	}
	v, _, _ := p.Call()
	w.Major = byte(v)
	w.Minor = uint8(v >> 8)
	w.Build = uint16(v >> 16)
	w.Version = fmt.Sprintf("%d.%d.%d", w.Major, w.Minor, w.Build)
	fmt.Println(fmt.Sprintf("[*] Detected version: %s", w.Version))
	return w
}

func GetHotfixesFromWMI(o *OutputConfig) []Hotfix {
	q := "SELECT HotFixID FROM Win32_QuickFixEngineering"
	o.OutputString("[+] Getting list of hotfixes from WMI")
	if !o.Quiet {
		o.OutputString(q)
	}
	var dst []Hotfix
	wmi.Query(q, &dst)
	return dst
}
