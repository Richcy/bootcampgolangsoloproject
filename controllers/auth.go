package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type AuthController struct {
	// declare variables
	store *session.Store
}

func InitAuthController(s *session.Store) *AuthController {
	return &AuthController{store: s}
}

// get /login
func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// post /login
func (controller *AuthController) LoginPosted(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	var myform LoginForm

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}
	// hardcode auth
	if myform.Username == "admin" && myform.Password == "1234" {
		sess.Set("username", "admin")
		sess.Save()

		return c.Redirect("/profile")
	}

	return c.Redirect("/login")
}

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
