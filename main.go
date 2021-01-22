package main

import (
	"evermos-online-store/config"
	"evermos-online-store/controller"
	"evermos-online-store/exception"
	"evermos-online-store/repository"
	"evermos-online-store/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	/*Setup Configuration*/
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	/*Setup Repository*/
	productRepository := repository.NewItemRepository(database)

	/*Setup Service*/
	itemService := service.NewItemService(&productRepository)

	/*Setup Controller*/
	productController := controller.NewItemController(&itemService)

	/*Setup Fiber*/
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	/*Setup Routing*/
	productController.Route(app)

	/*Start App*/
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
