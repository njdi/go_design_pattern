package strategy

const PremiumDiscount = 0.9

type DiscountStrategy interface {
	ApplyDiscount(price float32) float32
}

type RegularCustomerStrategy struct{}

func (r *RegularCustomerStrategy) ApplyDiscount(price float32) float32 {
	return price
}

type PremiumCustomerStrategy struct{}

func (p *PremiumCustomerStrategy) ApplyDiscount(price float32) float32 {
	return price * PremiumDiscount
}

type Context struct {
	strategy DiscountStrategy
}

func (c *Context) ExecuteStrategy(price float32) float32 {
	return c.strategy.ApplyDiscount(price)
}
