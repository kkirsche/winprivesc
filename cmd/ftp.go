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
	"github.com/kkirsche/winprivesc/libprivesc/helpers"
	"github.com/spf13/cobra"
)

var ftp helpers.FTP

// ftpCmd represents the ftp command
var ftpCmd = &cobra.Command{
	Use:   "ftp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ftp.Path = oc.FilePath
		ok := ftp.Setup()
		if ok {
			ftp.OutputFile()
		}
	},
}

func init() {
	helpersCmd.AddCommand(ftpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ftpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	ftpCmd.Flags().StringVarP(&ftp.IP, "ip-address", "i", "10.11.0.211", "IP address to connect to")
	ftpCmd.Flags().IntVarP(&ftp.Port, "port", "p", 21, "Port to connect to")

	ftpCmd.Flags().StringVarP(&ftp.Username, "username", "u", "d3c3pt10n", "FTP username")
	ftpCmd.Flags().StringVarP(&ftp.Password, "password", "s", "", "FTP password")

	ftpCmd.Flags().StringSliceVarP(&ftp.BinaryStrFiles, "binary", "b", []string{}, "Binary files to be transfered")
	ftpCmd.Flags().StringSliceVarP(&ftp.AsciiFilesStr, "ascii", "a", []string{}, "Ascii files to be transfered")
}
