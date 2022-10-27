package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"rapidtech/shoppingcart/database"
	"rapidtech/shoppingcart/models"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type CartController struct {
	// declare variables
	Db *gorm.DB
}

func InitCartController() *CartController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Cart{})

	return &CartController{Db: db}
}

// routing
// GET /products
func (controller *CartController) IndexCart(c *fiber.Ctx) error {
	// load all products
	var carts []models.Cart
	err := models.ReadCarts(controller.Db, &carts)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("carts", fiber.Map{
		"Title": "Cart",
		"Carts": carts,
	})
}

// GET /products/create
func (controller *CartController) AddCart(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	product := models.Product{}
	models.ReadProductById(controller.Db, &product, idn)
	fmt.Println(product)
	listproduk := []models.Product{}
	listproduk = append(listproduk, product)
	return c.Render("cart", fiber.Map{
		"Title": "Cart",
		"Carts": listproduk,
	})
}

// POST /products/create
func (controller *CartController) AddPostedCart(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var myform models.Cart

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}
	// save product

	err := models.CreateCart(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/products")
	}
	// if succeed
	return c.Redirect("/carts")
}

/*

// GET /products/productdetail?id=xxx
func (controller *ProductController) GetDetailProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

// GET /products/detail/xxx
func (controller *ProductController) GetDetailProduct2(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}
*/

// / GET products/editproduct/xx
/*
func (controller *CartController) EditCart(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var cart models.Cart
	err := models.ReadCartById(controller.Db, &cart, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("editcart", fiber.Map{
		"Title": "Edit Cart",
		"Cart":  cart,
	})
}


// / POST products/editproduct/xx
func (controller *CartController) EditPostedCart(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var cart models.Cart
	err := models.ReadCartById(controller.Db, &cart, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var myform models.Cart

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}
	cart.Name = myform.Name
	cart.Quantity = myform.Quantity
	cart.Price = myform.Price
	// save product
	models.UpdateCart(controller.Db, &cart)

	return c.Redirect("/products")

}
*/
// / GET /products/deleteproduct/xx
func (controller *CartController) DeleteCart(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var cart models.Cart
	models.DeleteCartById(controller.Db, &cart, idn)
	return c.Redirect("/products")
}
