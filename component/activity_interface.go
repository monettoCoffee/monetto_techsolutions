package component

import "context"

// ActivityComponentInterface 活动的抽象定义
type ActivityComponentInterface interface {
	Execute(ctx context.Context) error
}

// ActivityInfo 活动的数据对象 用于存放活动的具体数据
type ActivityInfo struct {
	// 活动ID
	ActivityID int64
	// 活动权益ID
	BenefitID int64
	// 订单ID
	OrderID string
}

// ActivityAction 定义触发活动的类型
type ActivityAction int64

const (
	ActivityAction_Unknown = 0
	// 通过下单参加活动
	ActivityAction_Order = 1
	// 订阅类型的活动
	ActivityAction_Subscribe = 2
)

// FetchUserCanJoinActivity 获取用户可以参加哪些活动
func FetchUserCanJoinActivity(activityAction ActivityAction) []ActivityComponentInterface {
	// TODO easy impl ...

	// 用户可以参加的活动列表
	activityComponentList := []ActivityComponentInterface{}
	if activityAction == ActivityAction_Order {
		activityComponentList = append(activityComponentList, OrderActivityComponent{})
	}

	return activityComponentList
}
