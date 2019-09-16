package policy

import (
	"github.com/pkg/errors"
	"github.com/txvier/dcoin/big"
	"github.com/txvier/dcoin/services"
	"strings"
)

//试算交易
type TrialTrade struct {
	Seq         int         //交易序列
	OrderPrice  big.Decimal //订单价格
	OrderAmount big.Decimal //订单数量
	//买入交易
	PurchaseAmount       big.Decimal  //买入数量
	PurchaseTrialAccount TrialAccount //交易类型
	//卖出交易
	SellOutAmount       big.Decimal  //卖出量
	SellOutTrialAccount TrialAccount //交易类型
}

// 试算策略
func TrialPolicy(p *Policy) (err error) {
	//p.Path:["usdt","eth","btc","gxc","usdt"]
	for _, step := range p.Steps {
		//step.Symbols: eth/usdt
		// 初始化策略步骤期初期末余额
		step.Purchase.InitialBalance = p.TrialAccounts[step.Purchase.Currency].FinalBalance
		step.Purchase.FinalBalance = p.TrialAccounts[step.Purchase.Currency].FinalBalance
		step.SellOut.InitialBalance = p.TrialAccounts[step.SellOut.Currency].FinalBalance
		step.SellOut.FinalBalance = p.TrialAccounts[step.SellOut.Currency].FinalBalance

		if err = TrialStep(step); err != nil {
			return err
		}
		p.TrialAccounts[step.Purchase.Currency].FinalBalance = step.Purchase.FinalBalance
		p.TrialAccounts[step.SellOut.Currency].FinalBalance = step.SellOut.FinalBalance
	}
	if p.Path[0] == p.Path[len(p.Path)-1] {
		bc := p.Path[0]
		p.Yields = p.TrialAccounts[bc].FinalBalance.Div(p.TrialAccounts[bc].InitialBalance)
	}
	return
}

// 试算某交易对步骤
func TrialStep(step *PolicyStep) (err error) {
	symbols := strings.Split(step.Symbols, "/")
	// 基础币种
	bc := symbols[0]
	// 计价币种
	qc := symbols[1]
	// 获取交易对市场深度数据
	md := services.GetMarketDepth(bc+qc, "150", "step0")
	if md.Status != "ok" {
		return errors.Errorf("get MarketDepth error for symbol:[%s]", bc+qc)
	}
	if bc == step.Purchase.Currency { //基础币种买入，看卖单
		for _, ask := range md.Tick.Asks {
			price, amount := big.NewDecimal(ask[0]), big.NewDecimal(ask[1])
			sellOutAmount := price.Mul(amount) //price * amount
			purchaseAmount := amount
			if step.SellOut.FinalBalance.GT(big.ZERO) && step.SellOut.FinalBalance.LT(sellOutAmount) {
				sellOutAmount = step.SellOut.FinalBalance
				purchaseAmount = sellOutAmount.Div(price)
			} else if step.SellOut.FinalBalance.LTE(big.ZERO) { // <=0
				break
			}
			createTrialTrade(step, amount, sellOutAmount, purchaseAmount, price)
		}
	} else { //基础币种卖出，看买单
		for _, bid := range md.Tick.Bids {
			price, amount := big.NewDecimal(bid[0]), big.NewDecimal(bid[1])
			sellOutAmount := amount
			purchaseAmount := price.Mul(amount) //price * amount
			if step.SellOut.FinalBalance.GT(big.ZERO) && step.SellOut.FinalBalance.LT(sellOutAmount) {
				sellOutAmount = step.SellOut.FinalBalance
				purchaseAmount = step.SellOut.FinalBalance.Mul(price)
			} else if step.SellOut.FinalBalance.LTE(big.ZERO) {
				break
			}
			createTrialTrade(step, amount, sellOutAmount, purchaseAmount, price)
		}
	}
	if step.SellOut.FinalBalance.GT(big.ZERO) {
		TrialStep(step)
	}
	return
}

func createTrialTrade(step *PolicyStep, orderAmount big.Decimal, sellOutAmount big.Decimal, purchaseAmount big.Decimal, price big.Decimal) {
	tt := newTrialTrace(step, orderAmount, sellOutAmount, purchaseAmount, price)
	tt.Seq = len(step.TrialTrades) + 1
	step.TrialTrades = append(step.TrialTrades, tt)
	step.Purchase.FinalBalance = tt.PurchaseTrialAccount.FinalBalance
	step.SellOut.FinalBalance = tt.SellOutTrialAccount.FinalBalance
}

func newTrialTrace(step *PolicyStep, orderAmount big.Decimal, sellOutAmount big.Decimal, purchaseAmount big.Decimal, price big.Decimal) (tt *TrialTrade) {
	tt = &TrialTrade{
		PurchaseAmount: purchaseAmount,
		OrderPrice:     price,
		OrderAmount:    orderAmount,
		PurchaseTrialAccount: TrialAccount{
			InitialBalance: step.Purchase.FinalBalance,
			Currency:       step.Purchase.Currency,
			FinalBalance:   step.Purchase.FinalBalance.Add(purchaseAmount),
		},
		SellOutAmount: sellOutAmount,
		SellOutTrialAccount: TrialAccount{
			InitialBalance: step.SellOut.FinalBalance,
			Currency:       step.SellOut.Currency,
			FinalBalance:   step.SellOut.FinalBalance.Sub(sellOutAmount),
		},
	}
	return
}
