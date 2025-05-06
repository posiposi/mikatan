package domain

import (
	"testing"
)

func TestStockValue(t *testing.T) {
	stock, _ := NewStock(true)
	if stock.Value() != true {
		t.Errorf("Value() returned %v, expected %v", stock.Value(), true)
	}

	stock, _ = NewStock(false)
	if stock.Value() != false {
		t.Errorf("Value() returned %v, expected %v", stock.Value(), false)
	}
}
