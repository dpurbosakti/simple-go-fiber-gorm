package handler

import (
	"fmt"

	"github.com/dpurbosakti/fiber-gorm/db"
	"github.com/dpurbosakti/fiber-gorm/models/user"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func GetAllUser(c *fiber.Ctx) error {
	users := []user.User{}
	if result := db.DB.Debug().Find(&users); result.Error != nil {
		log.Err(result.Error).Msg("failed to get all users data")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get all users data",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	user := user.User{}
	id := c.Params("id")
	if result := db.DB.Debug().Where("id = ?", id).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Err(result.Error).Msg("user not found")
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "failed to get data user",
				"error":   result.Error.Error(),
			})
		}
		log.Err(result.Error).Msg(fmt.Sprintf("failed to get data user with id %s", id))
		return result.Error
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	p := new(user.CreateUserRequest)

	if err := c.BodyParser(p); err != nil {
		log.Err(err).Msg("failed to parsing body")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to parsing body",
			"error":   err.Error(),
		})
	}

	reqData := user.ReqToEntity(p)
	if result := db.DB.Debug().Create(reqData); result.Error != nil {
		log.Err(result.Error).Msg("failed to create user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create user",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "created",
		"user":    reqData,
	})
}
