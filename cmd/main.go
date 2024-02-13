package main

import (
	"go-mart/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	// Define services
	inventoryService := services.NewInventoryService()

	app := fiber.New()

	app.Get("/inventory/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")
		log.Info("Getting inventory for store with id: " + id)

		inventory, err := inventoryService.GetInventoryByStoreId(id)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(inventory)
	})

	log.Fatal(app.Listen(":8080"))
}
