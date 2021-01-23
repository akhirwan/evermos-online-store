package validation

import (
	"evermos-online-store/exception"
	"evermos-online-store/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func CartValidate(request model.CreateCartRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ItemId, validation.Required),
		validation.Field(&request.ItemQty, validation.Required, validation.Min(0)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
