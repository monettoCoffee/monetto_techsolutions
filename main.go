package main

import (
	"context"
	"monettotechsolutions/param"
)

func main() {
	// 模拟客户端从手机下单 购买商品
	Order(context.Background(), &param.OrderRequest{
		OrderTimestamp: 0,
		UserID:         0,
		GoodItemID:     0,
		BranchShopID:   0,
	})

	// 模拟支付系统给业务系统回调 购买成功后 参加活动 下发权益
	OrderCallback(context.Background(), &param.OrderCallbackRequest{
		OrderID:        "",
		MerchantItemID: "",
		OrderSign:      "",
		OrderTimestamp: 0,
		UserID:         0,
		GoodItemID:     0,
		BranchShopID:   0,
		OrderStatus:    0,
	})

	// 模拟用户 通过 权益 ID 区分是享用免费的机会 还是 包月的机会
	CosumeBenefit(context.Background(), &param.ConsumeBenefitRequest{
		OrderTimestamp: 0,
		UserID:         0,
		BenefitID:      0,
	})

}
