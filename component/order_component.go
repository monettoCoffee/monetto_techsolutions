package component

import (
	"context"
	"encoding/json"
	"monettotechsolutions/rpc"
	"monettotechsolutions/utils"
)

// OrderComponent 订单的领域对象
type OrderComponent struct {
	// 生成的订单ID
	OrderID string
	// 商品在 App Store 或 Google Play 中的订阅商品号
	MerchantItemID int64
	// 下单签名
	OrderSign string
	// 下单幂等时间戳
	OrderTimestamp int64
	// 用户ID
	UserID int64
	// 商品ID
	GoodItemID int64
	// 在哪个门店下单的
	BranchShopID int64
	// 订单当前的状态
	OrderStatus OrderStatus
	// 订单签名
	Sign string
}

// OrderStatus 订单状态 主要表示系统内部状态
type OrderStatus int64

const (
	OrderStatus_UNKNOWN OrderStatus = 0
	// 下单成功
	OrderStatus_SUCCESS OrderStatus = 1
	// 下单失败
	OrderStatus_FAILED OrderStatus = 2
)

// Validate 校验对象是否合法 可以参与业务行为
func (h *OrderComponent) Validate() bool {
	// TODO easy impl ...

	if h.UserID == 0 {
		return false
	}

	return true
}

// Init 对领域对象进行初始化操作 例如 生成订单ID
func (h *OrderComponent) Init(ctx context.Context) error {
	// TODO easy impl ...

	// 生成订单 ID
	h.OrderID = utils.GenOrderID(ctx)

	// 生成预下单参数
	prepareOrderParam := &rpc.PrepareOrderParam{}

	// 向支付部门进行预下单 后续需要收取预下单的回调
	h.OrderSign = rpc.PrepareOrderRPC(ctx, prepareOrderParam)

	return nil
}

// SaveToDB 订单信息持久化到 DB 中
func (h *OrderComponent) SaveToDB() error {
	// TODO no impl ...
	return nil
}

// GetMerchantItemID 获取其在商店中的 真实商品ID 或者 订阅ID
func (h *OrderComponent) GetMerchantItemID() string {
	// TODO no impl ...
	return ""
}

// OrderComponentFromBusinessParam 从一个业务请求中 生成领域对象
func OrderComponentFromBusinessParam(ctx context.Context, value interface{}) *OrderComponent {
	// TODO easy impl ...
	jsonString := utils.ToLogStr(value)
	orderComponent := &OrderComponent{}
	json.Unmarshal([]byte(jsonString), orderComponent)

	// TODO lost error deal ...
	return orderComponent
}

// OrderComponentFromDatabase 从 DB 中 获取 DB 数据 生成领域对象
func OrderComponentFromDatabase(ctx context.Context, orderID string) *OrderComponent {
	// TODO get from mysql
	// TODO lost error deal ...
	return &OrderComponent{}
}
