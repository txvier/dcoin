package policy

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func initConfig() {
	viper.SetConfigFile("../dcoin.toml")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
}

func TestTrial(t *testing.T) {
	initConfig()
	per := []string{"usdt", "eth", "btc", "eos", "usdt"}
	Trial(per, "")
}

func TestCheckDuplicate(t *testing.T) {
	per := []string{"usdt", "eth", "btc", "gxc", "usdt"}
	if CheckDuplicate(per) {
		fmt.Println("duplicate...")
	}
}

func TestCreatePolicy(t *testing.T) {
	per := []string{"usdt", "eth", "btc", "gxc", "usdt"}
	CreatePolicy(per)
}

func TestTrialPolicy(t *testing.T) {
	per := []string{"usdt", "eth", "btc", "gxc", "usdt"}
	p, _ := CreatePolicy(per)
	TrialPolicy(p)
	PrintPolicy(*p)
}
