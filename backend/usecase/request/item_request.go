package request

type CreateItemRequest struct {
	ItemName    string
	Stock       bool
	Description string
	UserId      string
}

type UpdateItemRequest struct {
	ItemId      string
	ItemName    string
	Stock       bool
	Description string
}