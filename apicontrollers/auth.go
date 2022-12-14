package controllers

import (
	"rapidtech/shoppingcart/database"
	"rapidtech/shoppingcart/models"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"time"

	"gorm.io/gorm"
)

type LoginForm struct {
	Username string `form:"name" json:"name" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type AuthController struct {
	// declare variables
	//store *session.Store
	Db *gorm.DB
}

func InitAuthController() *AuthController {

	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.User{})
	return &AuthController{Db: db}
}

// POST /register
func (controller *AuthController) Register(c *fiber.Ctx) error {
	// load all user
	var myform models.User
	var convertpass LoginForm

	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(500)
	}
	convertpassword, _ := bcrypt.GenerateFromPassword([]byte(convertpass.Password), 10)
	sHash := string(convertpassword)

	myform.Password = sHash

	// save user
	err := models.CreateUser(controller.Db, &myform)
	if err != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": myform})
}

/*
// get /login
func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
*/

// POST /login
func (controller *AuthController) LoginPosted(c *fiber.Ctx) error {
	// load all products

	var user models.User
	var myform LoginForm
	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(500)
	}

	er := models.FindByUsername(controller.Db, &user, myform.Username)
	if er != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// hardcode auth
	mycompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
	if mycompare != nil {
		exp := time.Now().Add(time.Hour * 72)
		claims := jwt.MapClaims{
			"id":    user.Id,
			"name":  user.Name,
			"admin": true,
			"exp":   exp.Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("mysecretpassword"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"token":   t,
			"expired": exp.Format("2006-01-02 15:04:05"),
		})
	}
	return c.SendStatus(fiber.StatusUnauthorized)

}

/*
// /profile
func (controller *AuthController) Profile(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	val := sess.Get("username")

	return c.Render("myview", fiber.Map{
		"Title":    "Profile",
		"username": val,
	})
}

// /logout
func (controller *AuthController) Logout(c *fiber.Ctx) error {

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
*/
