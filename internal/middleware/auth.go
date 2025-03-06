package middleware

import (
	"crypto/sha256"
	"crypto/subtle"

	"fibertest/internal/config"
	"fibertest/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func Auth() func(*fiber.Ctx) error {
	cfg := config.GetConfig()
	apiKey := cfg.APIKey

	authMiddleware := keyauth.New(keyauth.Config{
		Validator: func(c *fiber.Ctx, key string) (bool, error) {
			hashedAPIKey := sha256.Sum256([]byte(apiKey))
			hashedKey := sha256.Sum256([]byte(key))

			if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
				return true, nil
			}
			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
		ErrorHandler: handler.ErrorHandler,
	})
	return authMiddleware
}
