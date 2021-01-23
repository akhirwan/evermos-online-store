package controller

import (
	"evermos-online-store/exception"
	"evermos-online-store/model"
	"evermos-online-store/service"

	"github.com/gofiber/fiber/v2"
)

type CartController struct {
	CartService service.CartService
}

func NewCartController(cartService *service.CartService) CartController {
	return CartController{CartService: *cartService}
}

func (controller *CartController) Route(app *fiber.App) {
	app.Post("/api/cart", controller.Create)
}

func (controller *CartController) Create(c *fiber.Ctx) error {
	var request model.CreateCartRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.CartService.Create(request, c.Get("user_email"))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
