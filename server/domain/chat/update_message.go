package chatdomain

import (
	"errors"
	"strings"
)

type UpdateMessageRequestBody struct {
	Message string `json:"message"`
}

func (body *UpdateMessageRequestBody) Validate() error {
	body.Message = strings.TrimSpace(body.Message)

	var err error
	if body.Message == "" {
		err = errors.Join(err, errors.New("message is empty"))
	}
	return err
}