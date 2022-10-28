package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"rapidtech/shoppingcart/database"
	"rapidtech/shoppingcart/models"
)

type TransactionController struct {
	// declare variables
	Db *gorm.DB
}

func InitTransactionController() *TransactionController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Transaction{})

	return &TransactionController{Db: db}
}

// routing
// GET /transactions
func (controller *TransactionController) IndexTransaction(c *fiber.Ctx) error {
	// load all products
	var transactions []models.Transaction
	err := models.ReadTransactions(controller.Db, &transactions)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transactions,
	})
}

// / DELETE /transaction/id
func (controller *TransactionController) DeleteTransactionById(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var transaction models.Transaction
	models.DeleteTransactionById(controller.Db, &transaction, idn)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transaction deleted",
	})
}
