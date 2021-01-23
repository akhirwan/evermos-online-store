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

type cartServiceImpl struct {
	CartRepository repository.CartRepository
	ItemRepository repository.ItemRepository
}

func NewCartService(cartRepository *repository.CartRepository, itemRepository *repository.ItemRepository) CartService {
	return &cartServiceImpl{
		CartRepository: *cartRepository,
		ItemRepository: *itemRepository,
	}
}

func (service *cartServiceImpl) Create(request model.CreateCartRequest, params ...string) (response model.CreateCartResponse) {
	request.Id = uuid.New().String()
	validation.CartValidate(request)

	/*check existing data*/
	findUser := service.CartRepository.FindAll(params[0])
	if findUser != nil {
		/*delete cart if exist*/
		service.Delete(findUser[0].Id)
	}

	findItem := service.ItemRepository.Show(request.ItemId)
	if findItem[0].IsDeleted == true {
		panic(exception.ValidationError{
			Message: "item is deleted",
		})
	}
	/*end check*/

	detailItem := model.DetailItems{
		ItemId:    request.ItemId,
		ItemName:  findItem[0].Name,
		ItemPrice: findItem[0].Price,
		ItemQty:   request.ItemQty,
	}

	cart := entity.Cart{
		Id:          request.Id,
		UserEmail:   params[0],
		CreatedAt:   time.Now().UnixNano() / int64(time.Millisecond),
		ModifiedAt:  time.Now().UnixNano() / int64(time.Millisecond),
		DetailItems: []interface{}{detailItem},
	}
	service.CartRepository.Insert(cart)

	response = model.CreateCartResponse{
		Id:          cart.Id,
		UserEmail:   cart.UserEmail,
		CreatedAt:   cart.CreatedAt,
		ModifiedAt:  cart.ModifiedAt,
		DetailItems: cart.DetailItems,
	}
	return response
}

func (service *cartServiceImpl) PutItem(request model.CreateCartRequest, params ...string) (response model.CreateCartResponse) {

	/*check existing data*/
	findCart := service.CartRepository.Show(params...)
	if findCart == nil {
		panic(exception.ValidationError{
			Message: "data not found",
		})
	}

	findItem := service.ItemRepository.Show(request.ItemId)
	if findItem[0].IsDeleted == true {
		panic(exception.ValidationError{
			Message: "item is deleted",
		})
	}
	/*end check*/

	detailItems := findCart[0].DetailItems

	detailItems = append(detailItems, entity.DetailItems{
		Id:       request.ItemId,
		Name:     findItem[0].Name,
		Price:    findItem[0].Price,
		Quantity: request.ItemQty,
	})

	cart := entity.Cart{
		Id:          request.Id,
		UserEmail:   params[0],
		CreatedAt:   time.Now().UnixNano() / int64(time.Millisecond),
		ModifiedAt:  time.Now().UnixNano() / int64(time.Millisecond),
		DetailItems: []interface{}{detailItems},
	}
	service.CartRepository.Update(cart)

	response = model.CreateCartResponse{
		Id:          cart.Id,
		UserEmail:   cart.UserEmail,
		CreatedAt:   cart.CreatedAt,
		ModifiedAt:  cart.ModifiedAt,
		DetailItems: cart.DetailItems,
	}

	return response
}

func (service *cartServiceImpl) Delete(params ...string) (response model.DeleteCartResponse) {

	/*check existing data*/
	findCart := service.CartRepository.Show(params...)
	if findCart == nil {
		panic(exception.ValidationError{
			Message: "data not found",
		})
	}
	/*end check*/

	cart := entity.Cart{Id: params[0]}
	service.CartRepository.Remove(cart)

	return response
}
