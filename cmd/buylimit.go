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
	//log "github.com/sirupsen/logrus"
	"github.com/txvier/dcoin/config"
	"github.com/txvier/dcoin/models"
	"github.com/txvier/dcoin/order"

	"github.com/spf13/cobra"
)

// buylimitCmd represents the buylimit command
var buylimitCmd = &cobra.Command{
	Use:   "buylimit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var p models.PlaceParams
		config.C.UnmarshalKey("order", &p)
		order.PlaceBuyLimit(p)
	},
}

func init() {
	orderCmd.AddCommand(buylimitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buylimitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buylimitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
