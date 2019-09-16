package base

import (
	"github.com/txvier/dcoin/big"
	"github.com/txvier/dcoin/config"
	"strings"
)

type Account struct {
	Currency string
	Balance  big.Decimal
}

type TradeSymbols struct {
	Symbol          string //交易对
	PricePrecision  int    // 报价精度
	AmountPrecision int    // 数量精度
}

var symbols map[string]*TradeSymbols

func GetTradeSymbol(s string) *TradeSymbols {
	key := strings.ReplaceAll(s, "/", "")
	return symbols[key]
}

type FuncGetAccounts func() (map[string]Account, error)
type FuncGetSymbols func() map[string]*TradeSymbols

var (
	GetAccounts FuncGetAccounts
	GetSymbols  FuncGetSymbols
)

func init() {
	container := initContainer()
	GetAccounts = container["GetAccounts_"+config.GetProfile()].(FuncGetAccounts)
	GetSymbols = container["GetSymbols_"+config.GetProfile()].(FuncGetSymbols)
}

func initContainer() map[string]interface{} {
	m := map[string]interface{}{
		"GetAccounts_dev":    GetAccountsDev,
		"GetAccounts_online": GetAccountsHuobi,
		"GetSymbols_dev":     GetSymbolsDev,
		"GetSymbols_online":  GetSymbolsHuobi,
		// add your interface here...
	}
	return m
}
