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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/txvier/dcoin/policy"
	"strings"
)

// trialCmd represents the trial command
var trialCmd = &cobra.Command{
	Use:   "trial",
	Short: "trial a policy",
	Long: `You can use "trial" command to trial a policy. e.g.:
			./dcoin trial`,
	Run: func(cmd *cobra.Command, args []string) {
		per := strings.Split(viper.GetString("policy.path"), " ")
		//policy.Trial(per, viper.GetString("policy.permutation"))
		policy.TrialTicker(per, viper.GetString("policy.permutation"))
	},
}

func init() {
	rootCmd.AddCommand(trialCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trialCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trialCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	trialCmd.Flags().String("policy", "", "--policy=\"usdt eth\"")

	viper.BindPFlag("policy", trialCmd.Flags().Lookup("policy"))
}
