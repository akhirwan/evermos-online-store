package model

type SubmitItemRequest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}

type SubmitItemResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int64  `json:"price"`
	Quantity   int32  `json:"quantity"`
	IsDeleted  bool   `json:"is_deleted"`
	CreatedAt  int64  `json:"created_at"`
	ModifiedAt int64  `json:"modified_at"`
}

type DeletedItemResponse struct {
	Id string `json:"id"`
}

type GetItemResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int64  `json:"price"`
	Quantity   int32  `json:"quantity"`
	IsDeleted  bool   `json:"is_deleted"`
	CreatedAt  int64  `json:"created_at"`
	ModifiedAt int64  `json:"modified_at"`
}
