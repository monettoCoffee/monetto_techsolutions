package handler

import (
	"context"
	"monettotechsolutions/component"
	"monettotechsolutions/param"
	"monettotechsolutions/utils"
)

// HandleOrder 处理支付下单的逻辑
func HandleOrder(ctx context.Context, request *param.OrderRequest) *param.OrderResponse {
	// 生成领域对象
	orderComponent := component.OrderComponentFromBusinessParam(ctx, request)

	// 业务初始化操作
	orderComponent.Init(ctx)

	// 领域对象的基本校验
	if !orderComponent.Validate() {
		utils.LogInfo(ctx, "[HandleOrder] Validate no passed ... %v", utils.ToLogStr(orderComponent))
		return nil
	}

	// 尝试保存到数据库中
	err := orderComponent.SaveToDB()
	if err != nil {
		utils.LogInfo(ctx, "[HandleOrder] Save to db failed %v, err: %v", utils.ToLogStr(orderComponent), err.Error())
		return nil
	}

	// 返回给客户端下单参数
	return &param.OrderResponse{
		OrderID:        orderComponent.OrderID,
		MerchantItemID: orderComponent.GetMerchantItemID(),
		OrderSign:      orderComponent.Sign,
	}
}

// HandleOrder 处理支付下单回调系统的回调逻辑
func HandleOrderCallback(ctx context.Context, request *param.OrderCallbackRequest) *param.OrderCallbackResponse {
	// 生成领域对象
	orderComponent := component.OrderComponentFromBusinessParam(ctx, request)

	// 从数据库中获取原始对象
	databaseOrder := component.OrderComponentFromDatabase(ctx, orderComponent.OrderID)

	// TODO 对比两者数据是否一致 做校验
	if databaseOrder != nil && databaseOrder.MerchantItemID != orderComponent.MerchantItemID {
		utils.LogInfo(ctx, "[HandleOrderCallback] Validate no passed ... %v, %v", utils.ToLogStr(orderComponent), utils.ToLogStr(databaseOrder))
		// TODO error deal ...
		return nil
	}

	// 保存到数据库
	orderComponent.SaveToDB()

	// TODO 这里如果需要做幂等 可以直接用 MQ 做 这里简单实现
	activityComponentList := component.FetchUserCanJoinActivity(component.ActivityAction_Order)
	if len(activityComponentList) == 0 {
		utils.LogInfo(ctx, "[HandleOrderCallback] No activity can join", utils.ToLogStr(orderComponent))
		return nil
	}

	// 依次对权益进行发放
	for _, activityAction := range activityComponentList {
		activityAction.Execute(ctx)
	}

	// 给支付系统发送回调完毕通知
	return &param.OrderCallbackResponse{}
}
