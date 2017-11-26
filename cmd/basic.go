// +build windows
// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/kkirsche/winprivesc/libprivesc"
	"github.com/spf13/cobra"
)

// basicCmd represents the basic command
var basicCmd = &cobra.Command{
	Use:   "basic",
	Short: "Run basic information gathering on the host",
	Long: `Execute basic windows privilege escalation information gathering. The
gathered data is based on FuzzySecurity's Windows Privilege Escalation
fundamentals.

URI: http://www.fuzzysecurity.com/tutorials/16.html`,
	Run: func(cmd *cobra.Command, args []string) {
		if oc.FilePath != "" {
			oc.FileEnabled = true
		}
		ok := oc.Setup()
		if ok {
			libprivesc.SystemChecks(&oc)
			libprivesc.NetworkChecks(&oc)
			libprivesc.Vulnerabilities(&oc)
		}
	},
}

func init() {
	RootCmd.AddCommand(basicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// basicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// basicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
