package libprivesc

import (
	"fmt"
	"os/exec"
)

// SystemChecks is used to run all system checks / commands
func SystemChecks() {
	fmt.Println("[+] Running system checks")
	SystemInfo()
	Users()
	fmt.Println("[+] System checks complete")
}

// SystemInfo will pull the system information
func SystemInfo() {
	fmt.Println("[+] Retrieving system information")
	out, err := exec.Command("cmd", "/c", "systeminfo").Output()
	if err != nil {
		fmt.Println("[!] Error retrieving system information")
		fmt.Println(fmt.Sprintf("\t%s", err))
	}
	fmt.Println(string(out))
}

// Users will pull the list of all user accounts
func Users() {
	fmt.Println("[+] Retrieving user list")
	out, err := exec.Command("cmd", "/c", "net users").Output()
	if err != nil {
		fmt.Println("[!] Error retrieving user list")
		fmt.Println(fmt.Sprintf("\t%s", err))
	}
	fmt.Println(string(out))
}
