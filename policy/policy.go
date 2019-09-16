package policy

import (
	"fmt"
	"github.com/apcera/termtables"
	"github.com/txvier/dcoin/base"
	"github.com/txvier/dcoin/big"
	"github.com/txvier/dcoin/permutation"
	"time"
)

// 试算账户
type TrialAccount struct {
	Currency       string      //币种
	InitialBalance big.Decimal //期初余额
	FinalBalance   big.Decimal //期末余额
}

//交易策略
type Policy struct {
	Path          []string                 //策略路径
	Steps         []*PolicyStep            //策略步骤
	Yields        big.Decimal              //收益率:其实币种的期末余额/期初余额
	TrialAccounts map[string]*TrialAccount //试算账户余额：每个币种对应的期初、期末余额
}

//基于排列创建策略
//per ["usdt","eth","btc","gxc","usdt"]
func CreatePolicy(per []string) (p *Policy, err error) {
	ats, err := initTrialAccounts(per)
	if err != nil {
		fmt.Println("initTrialAccounts error...")
		return
	}
	p = &Policy{
		Path:          per,
		TrialAccounts: ats,
		Yields:        big.ZERO,
	}
	//初始化策略步骤
	return p, InitPolicyStep(p)
}

func TrialTicker(path []string, prem string) {
	t := time.Tick(time.Second * 5)
	for range t {
		Trial(path, prem)
	}
}

func Trial(path []string, prem string) {
	if prem == "all" {
		ps, err := TrialAll(path)
		if err != nil {
			fmt.Println(err)
		}
		go func(ps []*Policy) {
			for _, p := range ps {
				PrintPolicy(*p)
			}
		}(ps)
	} else {
		p, err := TrialOne(path)
		if err != nil {
			fmt.Println(err)
		}
		go PrintPolicy(*p)
	}
}

func TrialOne(path []string) (p *Policy, err error) {
	if p, err = CreatePolicy(path); err != nil {
		return
	}
	return p, TrialPolicy(p)
}

func TrialAll(path []string) (ps []*Policy, err error) {
	pers := permutation.StartAsFlagEndEqBegin(path, path[0])
	for _, per := range pers {
		p, err := CreatePolicy(per)
		if err != nil {
			return nil, err
		}
		if err = TrialPolicy(p); err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return
}

// 初始化试算账户
func initTrialAccounts(per []string) (map[string]*TrialAccount, error) {
	ats, err := base.GetAccounts()
	if err != nil {
		return nil, err
	}
	m := slice2Map(per)
	for k := range m {
		if _, ok := ats[k]; !ok {
			m[k].InitialBalance = big.ZERO
			m[k].FinalBalance = big.ZERO
		} else {
			m[k].InitialBalance = ats[k].Balance
			m[k].FinalBalance = ats[k].Balance
		}
	}
	return m, nil
}

func slice2Map(per []string) map[string]*TrialAccount {
	if per[0] == per[len(per)-1] {
		per = per[:len(per)-1]
	}
	m := make(map[string]*TrialAccount)
	for _, v := range per {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = &TrialAccount{
			Currency: v,
		}
	}
	return m
}

func PrintPolicy(p Policy) {
	table := termtables.CreateTable()
	s := fmt.Sprint(p.Path, " | ", "收益率：", p.Yields)
	table.AddTitle(s)
	table.AddHeaders("Seq",
		"Symbols",
		"OrderPrice",
		"OrderAmount",
		"PurchaseAmount",
		"PC",
		"PurchaseInitialBalance",
		"PurchaseFinalBalance",
		"SellOutAmount",
		"SC",
		"SellOutInitialBalance",
		"SellOutFinalBalance",
	)
	for _, step := range p.Steps {
		pp := base.GetTradeSymbol(step.Symbols).PricePrecision
		ap := base.GetTradeSymbol(step.Symbols).AmountPrecision
		for _, tts := range step.TrialTrades {
			table.AddRow(tts.Seq,
				step.Symbols,
				tts.OrderPrice.FormattedString(pp),
				tts.OrderAmount.FormattedString(ap),
				tts.PurchaseAmount.FormattedString(ap),
				fmt.Sprint(tts.PurchaseTrialAccount.Currency),
				tts.PurchaseTrialAccount.InitialBalance.FormattedString(ap),
				tts.PurchaseTrialAccount.FinalBalance.FormattedString(ap),
				tts.SellOutAmount.FormattedString(ap),
				fmt.Sprint(tts.SellOutTrialAccount.Currency),
				tts.SellOutTrialAccount.InitialBalance.FormattedString(ap),
				tts.SellOutTrialAccount.FinalBalance.FormattedString(ap),
			)
		}
		table.AddSeparator()
	}
	fmt.Println(table.Render())
}
