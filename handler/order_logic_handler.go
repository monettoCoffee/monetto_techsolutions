package handler

import (
	"context"
	"monettotechsolutions/component"
	"monettotechsolutions/param"
)

// HandleBenifit 处理享受权益逻辑
func HandleBenifit(ctx context.Context, request *param.ConsumeBenefitRequest) *param.ConsumeBenefitResponse {
	// 生成领域对象
	benefitComponent := component.BenefitComponentFromDB(request.BenefitID)

	// 消费权益逻辑处理
	benefitComponent.ForUser()

	return &param.ConsumeBenefitResponse{}
}
