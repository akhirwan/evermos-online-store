package repository

import (
	"evermos-online-store/entity"
)

type CartRepository interface {
	FindAll(params ...string) (carts []entity.Cart)
	Show(params ...string) (carts []entity.Cart)
	Insert(cart entity.Cart)
	Update(cart entity.Cart)
	Remove(cart entity.Cart)
}
