package model

type GetCartResponse struct {
	Id          string
	UserEmail   string
	CreatedAt   int64
	ModifiedAt  int64
	DetailItems interface{}
}

type CreateCartRequest struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id"`
	ItemQty int32  `json:"item_qty"`
}

type CreateCartResponse struct {
	Id          string
	UserEmail   string
	CreatedAt   int64
	ModifiedAt  int64
	DetailItems interface{}
}

type DetailItems struct {
	ItemId    string
	ItemName  string
	ItemPrice int64
	ItemQty   int32
}

type DeleteCartResponse struct {
	Id string
}
