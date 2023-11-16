package main

import (
	"context"
	"monettotechsolutions/handler"
	"monettotechsolutions/param"
	"monettotechsolutions/utils"
)

// Order 通用下单接口 例如 来自于 iOS 客户端的下单
func Order(ctx context.Context, request *param.OrderRequest) *param.OrderResponse {
	utils.LogInfo(ctx, "[Order] Request: %v", utils.ToLogStr(request))
	resp := handler.HandleOrder(ctx, request)
	utils.LogInfo(ctx, "[Order] Response: %v", utils.ToLogStr(resp))
	return resp
}

// OrderCallback 用于接收支付系统回来的回调
func OrderCallback(ctx context.Context, request *param.OrderCallbackRequest) *param.OrderCallbackResponse {
	utils.LogInfo(ctx, "[OrderCallback] Request: %v", utils.ToLogStr(request))
	resp := handler.HandleOrderCallback(ctx, request)
	utils.LogInfo(ctx, "[OrderCallback] Response: %v", utils.ToLogStr(resp))
	return resp
}

// CosumeBenefit 用户请求享用权益
func CosumeBenefit(ctx context.Context, request *param.ConsumeBenefitRequest) *param.ConsumeBenefitResponse {
	utils.LogInfo(ctx, "[CosumeBenefit] Request: %v", utils.ToLogStr(request))
	resp := handler.HandleBenifit(ctx, request)
	utils.LogInfo(ctx, "[CosumeBenefit] Response: %v", utils.ToLogStr(resp))
	return resp
}
