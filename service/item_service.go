package service

import "evermos-online-store/model"

type ItemService interface {
	Create(request model.SubmitItemRequest) (response model.SubmitItemResponse)
	List(params ...string) (responses []model.GetItemResponse)
	Detail(params ...string) (responses []model.GetItemResponse)
	Edit(request model.SubmitItemRequest, params ...string) (response model.SubmitItemResponse)
	Remove(params ...string) (response model.DeletedItemResponse)
}
