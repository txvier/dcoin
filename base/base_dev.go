package base

import (
	"github.com/txvier/dcoin/big"
)

var GetSymbolsDev FuncGetSymbols = func() map[string]*TradeSymbols {
	if symbols != nil {
		return symbols
	}
	ds := []string{"btcusdt", "ethusdt", "ethbtc", "eosusdt", "eoseth", "eosbtc"}
	symbols = make(map[string]*TradeSymbols)
	for _, s := range ds {
		symbols[s] = &TradeSymbols{
			Symbol:          s,
			AmountPrecision: 8,
			PricePrecision:  8,
		}
	}
	return symbols
}

var GetAccountsDev FuncGetAccounts = func() (m map[string]Account, err error) {
	m = map[string]Account{
		"act": {
			Currency: "act",
			Balance:  big.NewDecimal(6513.4969000000),
		},
		"ela": {
			Currency: "ela",
			Balance:  big.NewDecimal(0.0000005800),
		},
		"eos": {
			Currency: "eos",
			Balance:  big.NewDecimal(150.2195982000),
		},
		"usdt": {
			Currency: "usdt",
			Balance:  big.NewDecimal(0.8397840080),
		},
		"eth": {
			Currency: "eth",
			Balance:  big.NewDecimal(48.1758641881),
		},
		"trx": {
			Currency: "trx",
			Balance:  big.NewDecimal(10),
		},
		"iost": {
			Currency: "iost",
			Balance:  big.NewDecimal(0.0054114000),
		},
		"aac": {
			Currency: "aac",
			Balance:  big.NewDecimal(10384.1800200000),
		},
		"btc": {
			Currency: "btc",
			Balance:  big.NewDecimal(1.8090025229),
		},
		"snt": {
			Currency: "snt",
			Balance:  big.NewDecimal(3403.4313962000),
		},
		"dbc": {
			Currency: "eth",
			Balance:  big.NewDecimal(15824.9367000000),
		},
		"dta": {
			Currency: "dta",
			Balance:  big.NewDecimal(0.0093046000),
		},
	}
	return
}
