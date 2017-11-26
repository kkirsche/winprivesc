// +build windows
package libprivesc

import (
	"fmt"
)

const patcherrmsg = "[!] Error checking for patch"

type Hotfix struct {
	HotFixID string
}

// Vulnerabilities runs all vulnerability checks
func Vulnerabilities(o *OutputConfig) {
	_ = GetWindowsVersion(o)
	hotfixes := GetHotfixesFromWMI(o)
	MS10015(hotfixes, o)
	MS10092(hotfixes, o)
	MS11080(hotfixes, o)
	MS13005(hotfixes, o)
	MS13053(hotfixes, o)
	MS13081(hotfixes, o)
}

// CheckForVuln is a helper command to execute a command via cmd /c
func (o *OutputConfig) CheckForVuln(premsg, kb, edb, ifVuln, errmsg string, hotfixes []Hotfix) {
	o.OutputString(premsg)
	if o.Verbose {
		o.OutputString(fmt.Sprintf("[*] Searching for KB: %s", kb))
	}
	found := false
	for _, hf := range hotfixes {
		if hf.HotFixID == kb {
			found = true
		}
	}

	if found {
		o.OutputString(fmt.Sprintf("[+] Patch %s found", kb))
	} else {
		o.OutputString(fmt.Sprintf("[$$] %s\n", ifVuln))
	}
}

// MS10015 (KB977165) exploit check
// https://www.exploit-db.com/exploits/11199/
func MS10015(hotfixes []Hotfix, o *OutputConfig) {
	premsg := "[+] Checking for MS10-015 patch (kitrap0d)"
	kb := "KB977165"
	edb := "https://www.exploit-db.com/exploits/11199/"
	ifVuln := "Possibly vulnerable to MS10-015 kitrap0d if Windows 2K SP4 - Windows 7 (x86)"
	errmsg := patcherrmsg
	o.CheckForVuln(premsg, kb, edb, ifVuln, errmsg, hotfixes)
}

// MS10092 (KB2305420) exploit check
// https://www.exploit-db.com/exploits/19930/
func MS10092(hotfixes []Hotfix, o *OutputConfig) {
	premsg := "[+] Checking for MS10-092 patch (schelevator)"
	kb := "KB2305420"
	edb := "https://www.exploit-db.com/exploits/19930/"
	ifVuln := "Possibly vulnerable to MS10-092 schelevator if Vista, 7, and 2008"
	errmsg := patcherrmsg
	o.CheckForVuln(premsg, kb, edb, ifVuln, errmsg, hotfixes)
}

// MS11080 (KB2592799) exploit check
// https://www.exploit-db.com/exploits/14610/
func MS11080(hotfixes []Hotfix, o *OutputConfig) {
	premsg := "[+] Checking for MS11-080 patch (afdjoinleaf)"
	edb := "https://www.exploit-db.com/exploits/18176/"
	kb := "KB2592799"
	ifVuln := "Possibly vulnerable to MS11-080 afdjoinleaf if XP SP2/SP3 Win 2k3 SP2"
	errmsg := patcherrmsg
	o.CheckForVuln(premsg, kb, edb, ifVuln, errmsg, hotfixes)
}

// MS13005 (KB2778930) exploit check
// https://www.exploit-db.com/exploits/14610/
func MS13005(hotfixes []Hotfix, o *OutputConfig) {
	premsg := "[+] Checking for MS13-005 patch (hwnd_broadcast)"
	edb := "https://www.exploit-db.com/exploits/24485/"
	kb := "KB2778930"
	ifVuln := "Possibly vulnerable to MS13-005 hwnd_broadcast, elevates from Low to Medium integrity"
	errmsg := patcherrmsg
	o.CheckForVuln(premsg, kb, edb, ifVuln, errmsg, hotfixes)
}

// MS13053 (KB2850851) exploit check
// https://www.exploit-db.com/exploits/33213/
func MS13053(hotfixes []Hotfix, o *OutputConfig) {
	premsg := "[+] Checking for MS13-053 patch (schlamperei)"
	edb := "https://www.exploit-db.com/exploits/33213/"
	kb := "KB2850851"
	ifVuln := "Possibly vulnerable to MS13-053 schlamperei if x86 Win7 SP0/SP1"
	errmsg := patcherrmsg
	o.CheckForVuln(premsg, kb, edb, ifVuln, errmsg, hotfixes)
}

// MS13081 (KB2870008) exploit check
// https://www.exploit-db.com/exploits/31576/
func MS13081(hotfixes []Hotfix, o *OutputConfig) {
	premsg := "[+] Checking for MS13-081 patch (track_popup_menu)"
	edb := "https://www.exploit-db.com/exploits/31576/"
	kb := "KB2870008"
	ifVuln := "Possibly vulnerable to MS13-081 track_popup_menu if x86 Windows 7 SP0/SP1"
	errmsg := patcherrmsg
	o.CheckForVuln(premsg, kb, edb, ifVuln, errmsg, hotfixes)
}
