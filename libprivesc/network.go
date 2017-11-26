// +build windows
package libprivesc

import (
	"strings"
)

// NetworkChecks is used to execute all network level comamnds / checks
func NetworkChecks(o *OutputConfig) {
	o.OutputString(strings.Repeat("#", 80))
	o.OutputString("# Network Checks")
	o.OutputString(strings.Repeat("#", 80))
	IPConfig(o)
	Routes(o)
	Arp(o)
	Netstat(o)
	FirewallState(o)
}

// IPConfig pulls the windows machine's IP data
func IPConfig(o *OutputConfig) {
	premsg := "[+] Retrieving ipconfig data"
	cmd := "ipconfig /all"
	errmsg := "[!] Error retrieving ipconfig data"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// Routes pulls the windows machine's route table
func Routes(o *OutputConfig) {
	premsg := "[+] Retrieving route data"
	cmd := "route print"
	errmsg := "[!] Error retrieving route data"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// Arp pulls the windows machine's arp data
func Arp(o *OutputConfig) {
	premsg := "[+] Retrieving arp data"
	cmd := "arp -A"
	errmsg := "[!] Error retrieving arp data"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// Netstat pulls the windows machine's netstat data
func Netstat(o *OutputConfig) {
	premsg := "[+] Retrieving netstat data"
	cmd := "netstat -ano"
	errmsg := "[!] Error retrieving netstat data"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// FirewallState pulls the windows machine's firewall state
func FirewallState(o *OutputConfig) {
	premsg := "[+] Retrieving firewall state"
	cmd := "netsh firewall show state"
	errmsg := "[!] Error retrieving firewall state"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// FirewallConfig pulls the windows machine's firewall config
func FirewallConfig(o *OutputConfig) {
	premsg := "[+] Retrieving firewall config"
	cmd := "netsh firewall show config"
	errmsg := "[!] Error retrieving firewall config"
	o.RunViaCmd(premsg, cmd, errmsg)
}
