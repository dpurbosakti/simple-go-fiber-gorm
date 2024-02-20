package handlers

import (
	"fmt"

	"github.com/dpurbosakti/fiber-gorm/app/models"
	"github.com/dpurbosakti/fiber-gorm/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (u *Handler) GetAllUser(c *fiber.Ctx) error {
	users, err := u.query.GetAllUser()

	if err != nil {
		log.Err(err).Msg("failed to get all users data")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get all users data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func (u *Handler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := u.query.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Err(err).Msg("user not found")
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "failed to get data user",
				"error":   err.Error,
			})
		}
		log.Err(err).Msg(fmt.Sprintf("failed to get data user with id %s", id))
		return err
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}

func (u *Handler) CreateUser(c *fiber.Ctx) error {
	p := new(models.CreateUserRequest)

	if err := c.BodyParser(p); err != nil {
		log.Err(err).Msg("failed to parsing body")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to parsing body",
			"error":   err.Error(),
		})
	}

	if err := u.validator.Struct(p); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	reqData := models.CreateReqToEntity(p)
	err := u.query.CreateUser(reqData)
	if err != nil {
		log.Err(err).Msg("failed to create user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create user",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "created",
	})
}

func (u *Handler) UpdateUser(c *fiber.Ctx) error {
	p := new(models.UpdateUserRequest)

	if err := c.BodyParser(p); err != nil {
		log.Err(err).Msg("failed to parsing body")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to parsing body",
			"error":   err.Error(),
		})
	}

	if err := u.validator.Struct(p); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	id := c.Params("id")
	user, err := u.query.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Err(err).Msg("user not found")
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "failed to get data user",
				"error":   err.Error,
			})
		}
		log.Err(err).Msg(fmt.Sprintf("failed to get data user with id %s", id))
		return err
	}

	models.UpdateReqToEntity(p, &user)
	err = u.query.UpdateUser(&user)
	if err != nil {
		log.Err(err).Msg("failed to update user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update user",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "updated",
	})
}
