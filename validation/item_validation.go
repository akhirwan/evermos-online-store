package validation

import (
	"evermos-online-store/exception"
	"evermos-online-store/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ItemValidate(request model.SubmitItemRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Price, validation.Required, validation.Min(1)),
		validation.Field(&request.Quantity, validation.Required, validation.Min(0)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
