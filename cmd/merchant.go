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

// merchantCmd represents the merchant command
var merchantCmd = &cobra.Command{
	Use:   "merchant",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("Invalid args for new merchant")
			return
		}
		value, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("Invalid discount value")
			return
		}
		err = service.GetMerchantService().SaveNewMerchant(args[0], args[1], value)
		if err != nil {
			fmt.Println("Rejected!")
			return
		}
		fmt.Printf("%s (%.2f)", args[0], value)
	},
}

func init() {
	newCmd.AddCommand(merchantCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// merchantCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// merchantCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
