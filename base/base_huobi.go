package base

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/txvier/dcoin/big"
	"github.com/txvier/dcoin/services"
	"strings"
)

var GetSymbolsHuobi FuncGetSymbols = func() map[string]*TradeSymbols {
	if symbols != nil {
		return symbols
	}
	sr := services.GetSymbols()
	if sr.Status != "ok" {
		fmt.Println("getSymbols from server error...")
		return nil
	}
	symbols = make(map[string]*TradeSymbols)
	for _, d := range sr.Data {
		symbols[d.Symbol] = &TradeSymbols{
			Symbol:          d.Symbol,
			AmountPrecision: d.AmountPrecision,
			PricePrecision:  d.PricePrecision,
		}
	}
	return symbols
}

var GetAccountsHuobi FuncGetAccounts = func() (m map[string]Account, err error) {
	ar := services.GetAccounts()
	if ar.Status != "ok" {
		return m, errors.New("get accounts error...")
	}
	logrus.Info("load accounts ok...")
	var aid string
	for _, d := range ar.Data {
		if d.Type == "spot" && d.State == "working" {
			aid = fmt.Sprint(d.ID)
		}
	}
	logrus.Infof("account spot aid is:[%s]\n", aid)
	br := services.GetAccountBalance(aid)
	if br.Status != "ok" {
		return m, errors.New("get accounts error...")
	}
	m = make(map[string]Account)
	for _, l := range br.Data.List {
		if l.Type == "trade" {
			if _, ok := m[l.Currency]; ok {
				continue
			}
			if strings.TrimSpace(l.Balance) != "0" {
				m[l.Currency] = Account{
					Currency: l.Currency,
					Balance:  big.NewFromString(l.Balance),
				}
			}
			logrus.Info("trade account balance is zero...\n")
		}
	}
	return
}
