package request

type CreateItemRequest struct {
	ItemName        string
	Stock           bool
	Description     string
	UserId          string
	PriceWithTax    *int
	PriceWithoutTax *int
	TaxRate         *float64
	Currency        *string
}

type UpdateItemRequest struct {
	ItemId          string
	ItemName        string
	Stock           bool
	Description     string
	PriceWithTax    *int
	PriceWithoutTax *int
	TaxRate         *float64
	Currency        *string
}