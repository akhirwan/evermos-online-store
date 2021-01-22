package service

import (
	"evermos-online-store/entity"
	"evermos-online-store/exception"
	"evermos-online-store/model"
	"evermos-online-store/repository"
	"evermos-online-store/validation"
	"time"

	"github.com/google/uuid"
)

type itemServiceImpl struct {
	ItemRepository repository.ItemRepository
}

func NewItemService(itemRepository *repository.ItemRepository) ItemService {
	return &itemServiceImpl{
		ItemRepository: *itemRepository,
	}
}
func (service *itemServiceImpl) Create(request model.SubmitItemRequest) (response model.SubmitItemResponse) {
	request.Id = uuid.New().String()
	validation.Validate(request)

	item := entity.Item{
		Id:         request.Id,
		Name:       request.Name,
		Price:      request.Price,
		Quantity:   request.Quantity,
		IsDeleted:  false,
		CreatedAt:  time.Now().UnixNano() / int64(time.Millisecond),
		ModifiedAt: time.Now().UnixNano() / int64(time.Millisecond),
	}

	service.ItemRepository.Insert(item)

	response = model.SubmitItemResponse{
		Id:         item.Id,
		Name:       item.Name,
		Price:      item.Price,
		Quantity:   item.Quantity,
		IsDeleted:  item.IsDeleted,
		CreatedAt:  item.CreatedAt,
		ModifiedAt: item.ModifiedAt,
	}

	return response
}

func (service *itemServiceImpl) List(params ...string) (responses []model.GetItemResponse) {
	products := service.ItemRepository.FindAll(params[0])
	for _, item := range products {
		responses = append(responses, model.GetItemResponse{
			Id:         item.Id,
			Name:       item.Name,
			Price:      item.Price,
			Quantity:   item.Quantity,
			IsDeleted:  item.IsDeleted,
			CreatedAt:  item.CreatedAt,
			ModifiedAt: item.ModifiedAt,
		})
	}

	return responses
}

func (service *itemServiceImpl) Detail(params ...string) (responses []model.GetItemResponse) {
	products := service.ItemRepository.Show(params[0])
	if products[0].IsDeleted == true {
		panic(exception.ValidationError{
			Message: "data is deleted",
		})
	}

	for _, item := range products {
		responses = append(responses, model.GetItemResponse{
			Id:         item.Id,
			Name:       item.Name,
			Price:      item.Price,
			Quantity:   item.Quantity,
			IsDeleted:  item.IsDeleted,
			CreatedAt:  item.CreatedAt,
			ModifiedAt: item.ModifiedAt,
		})
	}

	return responses
}

func (service *itemServiceImpl) Edit(request model.SubmitItemRequest, params ...string) (response model.SubmitItemResponse) {
	request.Id = params[0]
	validation.Validate(request)
	find := service.ItemRepository.Show(params[0])
	if find[0].IsDeleted == true {
		panic(exception.ValidationError{
			Message: "data is deleted",
		})
	}

	item := entity.Item{
		Id:         request.Id,
		Name:       request.Name,
		Price:      request.Price,
		Quantity:   request.Quantity,
		IsDeleted:  false,
		ModifiedAt: time.Now().UnixNano() / int64(time.Millisecond),
	}

	service.ItemRepository.Update(item)

	response = model.SubmitItemResponse{
		Id:         item.Id,
		Name:       item.Name,
		Price:      item.Price,
		Quantity:   item.Quantity,
		IsDeleted:  item.IsDeleted,
		CreatedAt:  find[0].CreatedAt,
		ModifiedAt: item.ModifiedAt,
	}

	return response
}

func (service *itemServiceImpl) Remove(params ...string) (response model.DeletedItemResponse) {

	find := service.ItemRepository.Show(params[0])
	if find[0].IsDeleted == true {
		panic(exception.ValidationError{
			Message: "data is deleted",
		})
	}

	item := entity.Item{
		Id:         params[0],
		IsDeleted:  true,
		ModifiedAt: time.Now().UnixNano() / int64(time.Millisecond),
	}

	service.ItemRepository.PutDelete(item)

	response = model.DeletedItemResponse{
		Id: item.Id,
	}

	return response
}
