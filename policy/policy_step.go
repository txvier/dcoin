package policy

import (
	"github.com/pkg/errors"
	"github.com/txvier/dcoin/base"
)

//策略步骤
type PolicyStep struct {
	StepNo      int           //步骤编号
	Symbols     string        //交易对
	Purchase    TrialAccount  //买入账户期初、期末余额
	SellOut     TrialAccount  //卖出账户期初、期末余额
	TrialTrades []*TrialTrade //试算交易
}

// 初始化策略步骤
func InitPolicyStep(p *Policy) (e error) {
	per := p.Path
	p.Steps = make([]*PolicyStep, len(per)-1)
	var pp *PolicyStep
	for i := 0; i < len(per); i++ {
		if i+1 == len(per) {
			break
		}
		// ["usdt","eth","btc","gxc","usdt"] -> ["eth/usdt","btc/eth","gxc/btc","usdt/gxc"]
		// 卖出币种
		sc := per[i]
		// 买入币种
		pc := per[i+1]
		// 实例化策略步骤
		pp = &PolicyStep{
			StepNo: i + 1,
			Purchase: TrialAccount{
				Currency: pc,
			},
			SellOut: TrialAccount{
				Currency: sc,
			},
		}
		// 初始化步骤交易对
		symbols := base.GetSymbols()
		s1, s2 := sc+pc, pc+sc
		if _, ok := symbols[s1]; !ok {
			//s1交易对没有，找s2
			if _, ok := symbols[s2]; !ok { //s2也没有返回错误
				return errors.Errorf("unSupport symbol:[%s,%s] for path:%v", s1, s2, per)
			} else { //s2有则
				pp.Symbols = pc + "/" + sc
			}
		} else { //s1有则
			pp.Symbols = sc + "/" + pc
		}
		// 新增策略步骤
		p.Steps[i] = pp
	}
	return
}
