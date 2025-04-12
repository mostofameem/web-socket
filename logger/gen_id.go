package logger

import (
	"log/slog"

	"github.com/google/uuid"
)

func GenerateID() string {
	ID, err := uuid.NewRandom()
	if err != nil {
		slog.Error("Failed to generate uuid", Extra(map[string]any{
			"error": err.Error(),
		}))
		return ""
	}

	return ID.String()
}
