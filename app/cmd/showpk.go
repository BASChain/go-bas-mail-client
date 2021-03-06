/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/BASChain/go-bas-mail-client/app/cmdclient"
	"github.com/BASChain/go-bas-mail-client/app/cmdcommon"
	"github.com/spf13/cobra"
	"log"
)

// showpkCmd represents the showpk command
var showpkCmd = &cobra.Command{
	Use:   "showpk",
	Short: "show public key",
	Long:  `show public key`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := cmdcommon.IsProcessStarted(); err != nil {
			log.Println(err)
			return
		}

		cmdclient.DefaultCmdSend("", cmdcommon.CMD_PK_SHOW)
	},
}

func init() {
	rootCmd.AddCommand(showpkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showpkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showpkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
