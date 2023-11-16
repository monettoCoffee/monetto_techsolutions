package param

// OrderRequest 通用的下单支付参数
type OrderRequest struct {
	// 下单幂等时间戳
	OrderTimestamp int64
	// 用户ID
	UserID int64
	// 商品ID
	GoodItemID int64
	// 在哪个门店下单的
	BranchShopID int64
}

// OrderResponse 下单响应参数
type OrderResponse struct {
	// 生成的 订单ID
	OrderID string
	// 商品在 App Store 或 Google Play 中的订阅商品号
	MerchantItemID string
	// 下单签名
	OrderSign string
}

// OrderCallbackRequest 支付回调的请求
type OrderCallbackRequest struct {
	// 生成的订单ID
	OrderID string
	// 商品在 App Store 或 Google Play 中的订阅商品号
	MerchantItemID string
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
	OrderStatus OrderStatusAPI
}

// OrderStatusAPI 订单状态
type OrderStatusAPI int64

const (
	OrderStatusAPI_UNKNOWN OrderStatusAPI = 0
	// 下单成功
	OrderStatusAPI_SUCCESS OrderStatusAPI = 1
	// 下单失败
	OrderStatusAPI_FAILED OrderStatusAPI = 2
)

// OrderCallbackResponse 下单之后 给支付系统的回调 主要是通知支付系统 RPC 回调已经接收到
type OrderCallbackResponse struct{}

// ConsumeBenefitRequest 用户享用权益的请求
type ConsumeBenefitRequest struct {
	// 下单幂等时间戳
	OrderTimestamp int64
	// 用户ID
	UserID int64
	// 权益ID
	BenefitID int64
}

// ConsumeBenefitResponse 响应用户享用权益的请求
type ConsumeBenefitResponse struct {
	// 成功 OR 失败 TODO easy impl ...
	ConsumeStatus int64
}
