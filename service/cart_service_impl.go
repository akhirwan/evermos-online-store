package service

import (
	"evermos-online-store/entity"
	"evermos-online-store/exception"
	"evermos-online-store/model"
	"evermos-online-store/repository"
	"evermos-online-store/validation"
	"log"
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

func (service *cartServiceImpl) Find(params ...string) (response []model.GetCartResponse) {
	log.Println(params[0])
	carts := service.CartRepository.Show(params[0])
	for _, cart := range carts {
		response = append(response, model.GetCartResponse{
			Id:          cart.Id,
			UserEmail:   cart.UserEmail,
			CreatedAt:   cart.CreatedAt,
			ModifiedAt:  cart.ModifiedAt,
			DetailItems: cart.DetailItems,
		})
	}
	return response
}

func (service *cartServiceImpl) Create(request model.CreateCartRequest, params ...string) (response model.CreateCartResponse) {
	request.Id = uuid.New().String()
	validation.CartValidate(request)

	/*check existing data*/
	findItem := service.ItemRepository.Show(request.ItemId)
	if findItem[0].IsDeleted == true {
		panic(exception.ValidationError{
			Message: "item is deleted",
		})
	}

	findUser := service.CartRepository.Show(params[0])
	/*end check*/

	detailItems := findUser[0].DetailItems

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
		DetailItems: detailItems,
	}

	if findUser != nil {
		service.CartRepository.Update(cart)
	} else {
		service.CartRepository.Insert(cart)
	}

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
		DetailItems: detailItems,
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
