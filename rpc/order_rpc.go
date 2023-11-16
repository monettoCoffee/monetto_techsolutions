package rpc

import (
	"context"
)

// PrepareOrderParam 支付侧预下单参数
type PrepareOrderParam struct {
}

// PrepareOrderRPC 准备向支付侧进行预下单
func PrepareOrderRPC(ctx context.Context, orderParam *PrepareOrderParam) string {
	// TODO easy impl ...
	return "SIGN"
}
