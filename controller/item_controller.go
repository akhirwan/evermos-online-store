package controller

import (
	"evermos-online-store/exception"
	"evermos-online-store/model"
	"evermos-online-store/service"

	"github.com/gofiber/fiber/v2"
)

type ItemController struct {
	ItemService service.ItemService
}

func NewItemController(itemService *service.ItemService) ItemController {
	return ItemController{ItemService: *itemService}
}

func (controller *ItemController) Route(app *fiber.App) {
	app.Post("/api/items", controller.Create)
	app.Get("/api/items", controller.List)
	app.Get("/api/items/:id", controller.Detail)
	app.Put("/api/items/:id", controller.Edit)
	app.Put("/api/items/:id/delete", controller.Delete)
}

func (controller *ItemController) Create(c *fiber.Ctx) error {
	var request model.SubmitItemRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.ItemService.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ItemController) List(c *fiber.Ctx) error {
	responses := controller.ItemService.List(c.Query("find"))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *ItemController) Detail(c *fiber.Ctx) error {
	responses := controller.ItemService.Detail(c.Params("id"))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *ItemController) Edit(c *fiber.Ctx) error {
	var request model.SubmitItemRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.ItemService.Edit(request, c.Params("id"))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ItemController) Delete(c *fiber.Ctx) error {
	responses := controller.ItemService.Remove(c.Params("id"))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
