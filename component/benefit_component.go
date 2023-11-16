package component

// BenefitComponent 权益类型领域对象 负责校验是否可以享受权益 以及 权益领域业务逻辑
type BenefitComponent struct {
	// 数据库中的权益 ID
	BenefitID int64
}

// SaveToDB 从 DB 中 查询相应的权益信息
func BenefitComponentFromDB(benefitID int64) *BenefitComponent {
	// TODO no impl ...
	return &BenefitComponent{}
}

// SaveToDB 权益信息持久化到 DB 中
func (h *BenefitComponent) SaveToDB() error {
	// TODO no impl ...
	return nil
}

// ForUser 将权益下发给用户
func (h *BenefitComponent) ForUser() error {
	// TODO no impl ...
	return nil
}

// Take 享受权益
func (h *BenefitComponent) Take() error {
	// TODO no impl ...
	return nil
}
