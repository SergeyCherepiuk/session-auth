package authhandler

import (
	"errors"
	"strings"
	"time"

	"github.com/SergeyCherepiuk/chat-app/logger"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (body LoginRequestBody) Validate() error {
	var err error

	if strings.TrimSpace(body.Username) == "" {
		err = errors.Join(errors.New("username is empty"))
	}

	if strings.TrimSpace(body.Password) == "" {
		err = errors.Join(err, errors.New("password is empty"))
	} else if len(body.Password) < 8 {
		err = errors.Join(err, errors.New("password is too short"))
	}

	return err
}

func (handler AuthHandler) Login(c *fiber.Ctx) error {
	body := LoginRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		logger.LogMessages <- logger.LogMessage{
			Message: "failed to parse the body",
			Level:   slog.LevelError,
			Attrs:   []slog.Attr{slog.String("err", err.Error())},
		}
		return err
	}

	if err := body.Validate(); err != nil {
		logger.LogMessages <- logger.LogMessage{
			Message: "request body isn't valid",
			Level:   slog.LevelError,
			Attrs: []slog.Attr{
				slog.String("err", err.Error()),
				slog.Any("body", body),
			},
		}
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	sessionId, userId, err := handler.storage.Login(body.Username, body.Password)
	if err != nil {
		logger.LogMessages <- logger.LogMessage{
			Message: "failed to log in user",
			Level:   slog.LevelError,
			Attrs:   []slog.Attr{slog.String("err", err.Error())},
		}
		return err
	}

	logger.LogMessages <- logger.LogMessage{
		Message: "user has been logged in",
		Level:   slog.LevelInfo,
		Attrs: []slog.Attr{
			slog.Uint64("user_id", uint64(userId)),
			slog.Any("session_id", sessionId),
		},
	}
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessionId.String(),
		HTTPOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})
	return c.SendStatus(fiber.StatusOK)
}