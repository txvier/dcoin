package order

import (
	"github.com/sirupsen/logrus"
	"github.com/txvier/dcoin/config"
	"github.com/txvier/dcoin/models"
	"github.com/txvier/dcoin/services"
	"strconv"
	"time"
)

func clientOrderIdGen() string {
	now := time.Now().UnixNano()
	sn := strconv.FormatInt(now, 10)
	// todo
	return sn
}

func place(p models.PlaceRequestParams) (r models.PlaceReturn) {
	paid := config.C.GetString(config.PLACE_ACCOUNT_ID)
	//checkAccountIdBalance(paid)
	//p.ClientOrderId = clientOrderIdGen()
	p.AccountID = paid
	p.Source = "api"
	r = services.Place(p)
	if r.Status != "ok" {
		// 下单失败
		logrus.Errorf("place order error[msg:%s]", r.ErrMsg)
	}
	return
}

// 限价买入
// 在当前最低价格下方挂买单（低价买入）
func PlaceBuyLimit(p models.PlaceParams) (r models.PlaceReturn) {
	p.Type = "buy-limit"
	prp := models.PlaceRequestParams{
		PlaceParams: p,
	}
	return place(prp)
}

// 限价买入-maker
// 在当前最低价格下方挂买单（低价买入）
func PlaceBuyLimitMaker(p models.PlaceParams) (r models.PlaceReturn) {
	p.Type = "buy-limit-maker"
	prp := models.PlaceRequestParams{
		PlaceParams: p,
	}
	return place(prp)
}

// 限价买入

// 在当前最高价格上方挂买单（高价买入）
func PlaceBuyStopLimit(p models.PlaceParams) (r models.PlaceReturn) {
	p.Type = "buy-stop-limit"
	p.Operator = "gte"
	prp := models.PlaceRequestParams{
		PlaceParams: p,
	}
	return place(prp)
}

// 限价卖出
// 在当前最高价格上方挂卖单（高价卖出）
func PlaceSellLimit(p models.PlaceParams) (r models.PlaceReturn) {
	p.Type = "sell-limit"
	prp := models.PlaceRequestParams{
		PlaceParams: p,
	}
	return place(prp)
}

// 限价卖出-maker
// 在当前最高价格上方挂卖单（高价卖出）
func PlaceSellLimitMaker(p models.PlaceParams) (r models.PlaceReturn) {
	p.Type = "sell-limit-maker"
	prp := models.PlaceRequestParams{
		PlaceParams: p,
	}
	return place(prp)
}

// 限价卖出
// 在当前最低价格下方挂卖单（低价卖出）
func PlaceSellStopLimit(p models.PlaceParams) (r models.PlaceReturn) {
	p.Type = "sell-stop-limit"
	p.Operator = "lte"
	prp := models.PlaceRequestParams{
		PlaceParams: p,
	}
	return place(prp)
}
