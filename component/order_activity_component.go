package component

import (
	"context"
	"monettotechsolutions/utils"
)

// OrderActivityComponent 活动的领域对象 面向订单类型活动
type OrderActivityComponent struct {
	// 活动的通用数据对象
	activityInfo ActivityInfo
}

// Execute 订单类型的活动逻辑 在这里下发权益
func (h OrderActivityComponent) Execute(ctx context.Context) error {
	utils.LogInfo(ctx, "[OrderActivityComponent] Execute by order ...")

	// TODO easy impl ...

	// 这里简单贴一下设计思路吧
	// 在权益表中 对某一种特定ID的权益 增加一条数据 状态为 未使用
	// 当用户需要使用这个权益的时候 对 这十条数据 变更状态为已使用
	// 而这个函数主要需要做的是 对权益表增加十条数据

	BenefitComponentFromDB(h.activityInfo.BenefitID).ForUser()

	return nil
}
