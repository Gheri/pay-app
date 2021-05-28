/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"pay-app/service"
	"strconv"

	"github.com/spf13/cobra"
)

// paybackCmd represents the payback command
var paybackCmd = &cobra.Command{
	Use:   "payback",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid args for user payback")
			return
		}
		value, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("Invalid value for payback")
		}
		err = service.GetUserService().Payback(args[0], value)
		if err != nil {
			fmt.Println(err)
			return
		}
		user, err := service.GetUserService().GetUser(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s(dues: %.2f)", user.Name, user.Dues)
	},
}

func init() {
	rootCmd.AddCommand(paybackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// paybackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// paybackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
