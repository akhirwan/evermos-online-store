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
	itemRepository := repository.NewItemRepository(database)
	cartRepository := repository.NewCartRepository(database)

	/*Setup Service*/
	itemService := service.NewItemService(&itemRepository)
	cartService := service.NewCartService(&cartRepository, &itemRepository)

	/*Setup Controller*/
	itemController := controller.NewItemController(&itemService)
	cartController := controller.NewCartController(&cartService)

	/*Setup Fiber*/
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	/*Setup Routing*/
	itemController.Route(app)
	cartController.Route(app)

	/*Start App*/
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
