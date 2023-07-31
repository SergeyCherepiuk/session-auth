package userhandler

import (
	"github.com/SergeyCherepiuk/chat-app/logger"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type GetUserResponseBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func (handler UserHandler) GetMe(c *fiber.Ctx) error {
	log := logger.Logger{}

	userId, _ := c.Locals("user_id").(uint)
	log.With(slog.Uint64("user_id", uint64(userId)))

	user, err := handler.storage.GetById(userId)
	if err != nil {
		log.Error("failed to get user by id", slog.String("err", err.Error()))
		return err
	}

	responseBody := GetUserResponseBody{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}
	log.Info("user's info has been sent to the user", slog.Any("user", responseBody))
	return c.JSON(responseBody)
}
