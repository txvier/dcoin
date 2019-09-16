/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"github.com/txvier/dcoin/base"
	"strings"

	"github.com/spf13/cobra"
)

// symbolsCmd represents the symbols command
var listSymbolsCmd = &cobra.Command{
	Use:   "symbols",
	Short: "List symbols include the args",
	Args:  cobra.MinimumNArgs(1),
	Long:  `List symbols include the args`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			s := base.GetSymbols()
			for k, v := range s {
				if strings.Contains(k, arg) {
					fmt.Printf("Symbol[%s] - PricePrecision[%d] - AmountPrecision[%d]\n", v.Symbol, v.PricePrecision, v.AmountPrecision)
				}
			}
		}
	},
}

func init() {
	listCmd.AddCommand(listSymbolsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// symbolsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// symbolsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
