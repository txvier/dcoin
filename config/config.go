package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/txvier/dcoin/base"
	"os"
)

const (
	CFG_FILE         = "dcoin.toml"
	PLACE_ACCOUNT_ID = "place_account_id"
)

var C *viper.Viper

var profile string

func GetProfile() string {
	if profile != "" {
		return profile
	}
	profile = C.GetString("profile")
	return profile
}

func init() {

	if CFG_FILE != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CFG_FILE)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".dcoin" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dcoin")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println()
		fmt.Println("Config file changed:", e.Name)
		base.PrintDcoinPrefix()
	})

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		C = viper.GetViper()
	}
}
