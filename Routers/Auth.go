package routers

import (
	"FiberAuthWithSurrealDb/Database"
	"FiberAuthWithSurrealDb/Util"
	"github.com/gofiber/fiber/v2"
)

type Body struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Login(c *fiber.Ctx) error {

	body := new(Body)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	if body.Username != "" && body.Password != "" {
		token, err := database.VerifyUser(body.Username, body.Password)

		if err != nil {

			if err.Error() == "user_not_found" {
				return c.JSON(fiber.Map{
					"error": "User not found",
				})
			} else if err.Error() == "invalid_password" {
				return c.JSON(fiber.Map{
					"error": "Invalid password",
				})
			} else {
				return c.JSON(fiber.Map{
					"error": "Something went wrong",
				})
			}
		}

		if token != nil {

			return c.JSON(fiber.Map{
				"message": "Logged in",
				"token":   token,
			})
		} else {
			return c.JSON(fiber.Map{
				"error": "Loggin failed",
			})
		}
	} else {
		return c.JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}
}

func Register(c *fiber.Ctx) error {
	body := new(Body)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	if body.Username != "" && body.Password != "" && body.Email != "" {
		if util.ValidateUsername(body.Username) && util.ValidatePassword(body.Password) && util.ValidateEmail(body.Email) {
			_, err := database.CreateUser(body.Username, body.Password, body.Email)

			if err != nil {
				if err.Error() == "user_already_exists" {
					return c.JSON(fiber.Map{
						"message": "User already exists",
					})
				} else if err.Error() == "email_already_exists" {
					return c.JSON(fiber.Map{
						"message": "Email already exists",
					})
				} else {
					return c.JSON(fiber.Map{
						"message": "User not created",
					})
				}
			} else {
				token := util.SignToken(util.JWTData{
					Username: body.Username,
					Email:    body.Email,
					Password: body.Password,
				})

				return c.JSON(fiber.Map{
					"message": "User created",
					"token":   token,
				})
			}
		}
	} else {
		return c.JSON(fiber.Map{
			"message": "User not created",
		})
	}
	return nil
}
