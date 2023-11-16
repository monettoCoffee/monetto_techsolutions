package component

import (
	"context"
	"monettotechsolutions/utils"
)

// SubscribeActivityComponent 订阅的领域对象 面向订阅类型活动
type SubscribeActivityComponent struct {
	// 活动的通用数据对象
	activityInfo ActivityInfo
}

// Execute 订单类型的活动逻辑 在这里下发权益
func (h SubscribeActivityComponent) Execute(ctx context.Context) error {
	utils.LogInfo(ctx, "[SubscribeActivityComponent] Execute by subscribe ...")

	// 设计思路为 直接创建一个权益类型对象 保存到 DB 中
	// 用户想要享受无限次权益的时候 直接从接口调用获取
	// TODO easy impl ...

	BenefitComponentFromDB(h.activityInfo.BenefitID).ForUser()

	return nil
}
