package huobi

import (
	"fmt"
	"github.com/txvier/dcoin/services"
)

type HBSymbolsConv struct {
	Symbols map[string]*TradeSymbols
}

type TradeSymbols struct {
	//TradeDirection string //交易方向：买:"1";卖:"-1"
	BaseCurrency  string // 基础币种
	QuoteCurrency string // 计价币种
	Symbol        string // 交易对
}

func (d HBSymbolsConv) GetSymbols() map[string]*TradeSymbols {
	return d.Symbols
}

func (d HBSymbolsConv) ConvSymbol(bc, qc, symbol string) (bs *TradeSymbols, err error) {
	/*s1, s2 := bc+qc, qc+bc
	if _, ok := d.Symbols[s1]; !ok {
		//s1交易对没有，找s2
		if _, ok := d.Symbols[s2]; !ok {
			//s2也没有返回错误
			return bs, errors.Errorf("unsupport symbol[%s,%s]", s1, s2)
		}
		//s2有则
		bs = d.Symbols[s2]
		bs.TradeDirection = TRADE_DIRECTION_P
		return
	}
	//s1有则
	bs = d.Symbols[s1]
	bs.TradeDirection = TRADE_DIRECTION_S*/
	return
}

func GetSymbolsConv() SymbolsConv {
	sr := services.GetSymbols()
	if sr.Status != "ok" {
		fmt.Println("getSymbols fail...")
	}
	sc := HBSymbolsConv{
		Symbols: map[string]*TradeSymbols{},
	}
	for _, d := range sr.Data {
		sc.Symbols[d.Symbol] = &TradeSymbols{
			Symbol:        d.Symbol,
			BaseCurrency:  d.BaseCurrency,
			QuoteCurrency: d.QuoteCurrency,
		}
	}
	return sc
}

type SymbolsConv interface {
	ConvSymbol(bc, qc, symbol string) (bs *TradeSymbols, err error)
	GetSymbols() map[string]*TradeSymbols
}
