package models

type SymbolsData struct {
	BaseCurrency    string  `json:"base-currency"`    // 基础币种
	QuoteCurrency   string  `json:"quote-currency"`   // 计价币种
	Symbol          string  `json:"symbol"`           // 交易对
	PricePrecision  int     `json:"price-precision"`  // 价格精度位数(0为个位)
	AmountPrecision int     `json:"amount-precision"` // 数量精度位数(0为个位)
	SymbolPartition string  `json:"symbol-partition"` // 交易区, main: 主区, innovation: 创新区, bifurcation: 分叉区
	State           string  `json:"state"`            // 交易对状态；可能值: [online，offline,suspend] online - 已上线；offline - 交易对已下线，不可交易；suspend -- 交易暂停
	MinOrderAmt     float64 `json:"min-order-amt"`    //交易对最小下单量 (下单量指当订单类型为限价单或sell-market时，下单接口传的'amount')
	MaxOrderAmt     float64 `json:"max-order-amt"`    //交易对最大下单量
	MinOrderValue   float64 `json:"min-order-value"`  //最小下单金额 （下单金额指当订单类型为限价单时，下单接口传入的(amount * price)。当订单类型为buy-market时，下单接口传的'amount'）
}

type SymbolsReturn struct {
	Status  string        `json:"status"` // 请求状态
	Data    []SymbolsData `json:"data"`   // 交易及精度数据
	ErrCode string        `json:"err-code"`
	ErrMsg  string        `json:"err-msg"`
}
