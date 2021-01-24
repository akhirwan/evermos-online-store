package service

import (
	"evermos-online-store/model"
)

type CartService interface {
	Find(params ...string) (response []model.GetCartResponse)
	Create(request model.CreateCartRequest, params ...string) (response model.CreateCartResponse)
	PutItem(request model.CreateCartRequest, params ...string) (response model.CreateCartResponse)
	Delete(params ...string) (response model.DeleteCartResponse)
}
