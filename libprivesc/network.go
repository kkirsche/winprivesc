package libprivesc

import (
	"fmt"
	"os/exec"
)

// NetworkChecks is used to execute all network level comamnds / checks
func NetworkChecks() {
	fmt.Println("[+] Running network checks")
	IPConfig()
	Routes()
	Arp()
	Netstat()
	fmt.Println("[+] Network checks complete")
}

// IPConfig pulls the windows machine's IP data
func IPConfig() {
	fmt.Println("[+] Retrieving ipconfig /all data")
	out, err := exec.Command("cmd", "/c", "ipconfig /all").Output()
	if err != nil {
		fmt.Println("[!] Error retrieving ipconfig data")
		fmt.Println(fmt.Sprintf("\t%s", err))
	}
	fmt.Println(string(out))
}

// Routes pulls the windows machine's route table
func Routes() {
	fmt.Println("[+] Retrieving route data")
	out, err := exec.Command("cmd", "/c", "route print").Output()
	if err != nil {
		fmt.Println("[!] Error retrieving route data")
		fmt.Println(fmt.Sprintf("\t%s", err))
	}
	fmt.Println(string(out))
}

// Arp pulls the windows machine's arp data
func Arp() {
	fmt.Println("[+] Retrieving arp data")
	out, err := exec.Command("cmd", "/c", "arp -A").Output()
	if err != nil {
		fmt.Println("[!] Error retrieving arp data")
		fmt.Println(fmt.Sprintf("\t%s", err))
	}
	fmt.Println(string(out))
}

// Netstat pulls the windows machine's netstat data
func Netstat() {
	fmt.Println("[+] Retrieving netstat data")
	out, err := exec.Command("cmd", "/c", "netstat -ano").Output()
	if err != nil {
		fmt.Println("[!] Error retrieving netstat data")
		fmt.Println(fmt.Sprintf("\t%s", err))
	}
	fmt.Println(string(out))
}
