// +build windows
package libprivesc

import (
	"strings"
)

// SystemChecks is used to run all system checks / commands
func SystemChecks(o *OutputConfig) {
	o.OutputString(strings.Repeat("#", 80))
	o.OutputString("# System Checks")
	o.OutputString(strings.Repeat("#", 80))
	SystemInfo(o)
	Users(o)
	ScheduledTasks(o)
	ServiceTasklist(o)
	StartedServices(o)
	Drivers(o)
}

// SystemInfo will pull the system information
func SystemInfo(o *OutputConfig) {
	premsg := "[+] Retrieving system information"
	cmd := "systeminfo"
	errmsg := "[!] Error retrieving system information"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// Users will pull the list of all user accounts
func Users(o *OutputConfig) {
	premsg := "[+] Retrieving user list"
	cmd := "net users"
	errmsg := "[!] Error retrieving user list"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// ScheduledTasks will pull the list of scheduled tasks from the machine
func ScheduledTasks(o *OutputConfig) {
	premsg := "[+] Retrieving scheduled tasks"
	cmd := "schtasks /query /fo LIST /v"
	errmsg := "[!] Error retrieving scheduled tasks"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// ServiceTasklist will pull the list of service tasklist from the machine
func ServiceTasklist(o *OutputConfig) {
	premsg := "[+] Retrieving service tasklist"
	cmd := "tasklist /SVC"
	errmsg := "[!] Error retrieving service tasklist"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// StartedServices lists the started windows services
func StartedServices(o *OutputConfig) {
	premsg := "[+] Retrieving started services"
	cmd := "net start"
	errmsg := "[!] Error retrieving started services"
	o.RunViaCmd(premsg, cmd, errmsg)
}

// Drivers lists the drivers on the machine
func Drivers(o *OutputConfig) {
	premsg := "[+] Retrieving drivers"
	cmd := "DRIVERQUERY"
	errmsg := "[!] Error retrieving drivers"
	o.RunViaCmd(premsg, cmd, errmsg)
}
