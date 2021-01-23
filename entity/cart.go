package entity

type Cart struct {
	Id          string
	UserEmail   string
	CreatedAt   int64
	ModifiedAt  int64
	DetailItems []interface{} /*call item entity*/
}

type DetailItems struct {
	Id       string
	Name     string
	Price    int64
	Quantity int32
}
