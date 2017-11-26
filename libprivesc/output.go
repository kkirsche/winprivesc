// +build windows
package libprivesc

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// OutputConfig is leveraged to ensure we output as the user requested
type OutputConfig struct {
	FileEnabled bool
	FilePath    string
	Quiet       bool
	Verbose     bool
}

// RunDirect is a helper command to execute a command directly without calling
// cmd /c first
func (o *OutputConfig) RunDirect(premsg, cmd, errmsg string) {
	o.OutputString(premsg)
	if o.Verbose {
		o.OutputString(fmt.Sprintf("[*] Executing Command: %s", cmd))
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		o.OutputString(errmsg)
		o.OutputString(fmt.Sprintf("\t%s", err))
	}
	o.OutputBytes(out)
}

// RunViaCmd is a helper command to execute a command via cmd /c
func (o *OutputConfig) RunViaCmd(premsg, cmd, errmsg string) {
	o.OutputString(premsg)
	if o.Verbose {
		o.OutputString(fmt.Sprintf("[*] Executing Command: cmd /c %s", cmd))
	}
	out, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		o.OutputString(errmsg)
		o.OutputString(fmt.Sprintf("\t%s", err))
	}
	o.OutputBytes(out)
}

// Setup is used to prepare the output
func (o OutputConfig) Setup() bool {
	if o.FileEnabled {
		f, err := os.OpenFile(o.FilePath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("[!] Failed to open file for report")
			fmt.Println(fmt.Sprintf("\t%s", err))
			return false
		}
		defer f.Close()
		f.WriteString("Windows Privilege Escalation Check Report\n")
		f.WriteString("v1.0.0 | by Kevin Kirsche / d3c3pt10n\n")
		f.WriteString(fmt.Sprintf("Write to file: %t | Path: %s\n", o.FileEnabled, o.FilePath))
		f.WriteString(fmt.Sprintf("Verbose: %t\n", o.Verbose))
		f.WriteString(fmt.Sprintf("Quiet: %t\n", o.Quiet))
		f.WriteString("\n")
	}

	if !o.Quiet {
		fmt.Println(strings.Repeat("#", 80))
		fmt.Println("Windows Privilege Escalation Check Report")
		fmt.Println("v1.0.0 | by Kevin Kirsche / d3c3pt10n")
		fmt.Println(fmt.Sprintf("Write to file: %t | Path: %s", o.FileEnabled, o.FilePath))
		fmt.Println(fmt.Sprintf("Verbose: %t", o.Verbose))
		fmt.Println(fmt.Sprintf("Quiet: %t", o.Quiet))
		fmt.Println(strings.Repeat("#", 80))
		fmt.Println("")
	}

	time.Sleep(5 * time.Second)

	return true
}

// OutputString is used to output string data to a file or stdout
func (o *OutputConfig) OutputString(line string) {
	if o.FileEnabled {
		f, err := os.OpenFile(o.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			fmt.Println("[!] Failed to open file for report")
			fmt.Println(fmt.Sprintf("\t%s", err))
		}
		defer f.Close()
		f.WriteString(line)
		f.WriteString("\n")
	}

	if !o.Quiet {
		fmt.Println(line)
	}
}

// OutputBytes is used to output byte data to a file or stdout
func (o *OutputConfig) OutputBytes(line []byte) {
	if o.FileEnabled {
		f, err := os.OpenFile(o.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			fmt.Println("[!] Failed to open file for report")
			fmt.Println(fmt.Sprintf("\t%s", err))
		}
		defer f.Close()
		f.Write(line)
		f.WriteString("\n")
	}

	if !o.Quiet {
		fmt.Println(string(line))
	}
}
