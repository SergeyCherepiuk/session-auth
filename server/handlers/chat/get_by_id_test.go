package chathandler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TestUnauthorizedGetByIdRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/1", nil)

	response, _ := app.Test(request)
	if response.StatusCode != fiber.StatusUnauthorized {
		t.Errorf(
			"expected status code: %v, actual status code: %v\n",
			fiber.StatusUnauthorized,
			response.StatusCode,
		)
	}
}

func TestInvalidGetByIdRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/asd", nil)
	request.AddCookie(&http.Cookie{
		Name:     "session_id",
		Value:    uuid.NewString(),
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	response, _ := app.Test(request)
	if response.StatusCode != fiber.StatusBadRequest {
		t.Errorf(
			"expected status code: %v, actual status code: %v\n",
			fiber.StatusBadRequest,
			response.StatusCode,
		)
	}
}

func TestChatNotFoundGetByIdRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/2", nil)
	request.AddCookie(&http.Cookie{
		Name:     "session_id",
		Value:    uuid.NewString(),
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	response, _ := app.Test(request)
	if response.StatusCode != fiber.StatusInternalServerError {
		t.Errorf(
			"expected status code: %v, actual status code: %v\n",
			fiber.StatusInternalServerError,
			response.StatusCode,
		)
	}
}

func TestValidGetByIdRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/1", nil)
	request.AddCookie(&http.Cookie{
		Name:     "session_id",
		Value:    uuid.NewString(),
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	response, _ := app.Test(request)
	if response.StatusCode != fiber.StatusOK {
		t.Errorf(
			"expected status code: %v, actual status code: %v\n",
			fiber.StatusOK,
			response.StatusCode,
		)
	}
}