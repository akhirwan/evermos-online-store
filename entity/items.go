package entity

type Item struct {
	Id         string
	Name       string
	Price      int64
	Quantity   int32
	IsDeleted  bool
	CreatedAt  int64
	ModifiedAt int64
}
