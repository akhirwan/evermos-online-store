package repository

import (
	"evermos-online-store/entity"
)

type ItemRepository interface {
	Insert(item entity.Item)

	FindAll(params ...string) (items []entity.Item)
	Show(params ...string) (items []entity.Item)

	Update(item entity.Item)
	PutDelete(item entity.Item)
}
