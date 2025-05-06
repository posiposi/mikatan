package domain

type Stock struct {
	value bool
}

func NewStock(value bool) (*Stock, error) {
	stock := new(Stock)
	stock.value = value
	return stock, nil
}

func (stock Stock) Value() bool {
	return stock.value
}
