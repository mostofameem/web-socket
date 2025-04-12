package utils

import (
	"log/slog"
	"net/url"
	"websocket/logger"

	"github.com/go-playground/form/v4"
)

func BindValues(v interface{}, values url.Values) error {
	dec := form.NewDecoder()
	if err := dec.Decode(v, values); err != nil {
		slog.Error("Failed to bind values", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}

	return nil
}
