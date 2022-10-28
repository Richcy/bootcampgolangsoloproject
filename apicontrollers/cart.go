package controllers

import (
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
// GET /carts
func (controller *CartController) IndexCart(c *fiber.Ctx) error {
	// load all products
	var carts []models.Cart
	err := models.ReadCarts(controller.Db, &carts)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": carts})
}

// POST /products/addcart/xx
func (controller *CartController) AddPostedCart(c *fiber.Ctx) error {
	//myform := new(models.Product)
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	var myform models.Cart

	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(400)
	}

	myform.Id = product.Id
	myform.Name = product.Name
	myform.Image = product.Image
	myform.Price = product.Price
	// save cart
	errr := models.CreateCart(controller.Db, &myform)
	if errr != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.JSON(myform)
}

/*
func (controller *CartController) AddCart(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	product := models.Product{}
	models.ReadProductById(controller.Db, &product, idn)
	fmt.Println(product)
	listproduk := []models.Product{}
	listproduk = append(listproduk, product)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": listproduk})
}
*/

// POST /carts/checkout/xx
func (controller *CartController) Checkout(c *fiber.Ctx) error {
	//myform := new(models.Product)
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	var cart models.Cart
	err := models.ReadCartById(controller.Db, &cart, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	var myform models.Transaction

	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(400)
	}

	myform.Id = cart.Id
	myform.Name = cart.Name
	myform.Image = cart.Image
	myform.Price = cart.Price
	// add to transaction
	errr := models.CreateTransaction(controller.Db, &myform)
	if errr != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.JSON(myform)
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
// / DELETE /carts/deleteproduct/xx
func (controller *CartController) DeleteCartById(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var cart models.Cart
	models.DeleteCartById(controller.Db, &cart, idn)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "product in cart deleted",
	})
}

// / DELETE /carts/deletecart
func (controller *CartController) DeleteCart(c *fiber.Ctx) error {

	var cart models.Cart
	models.DeleteCart(controller.Db, &cart)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "cart deleted",
	})
}
