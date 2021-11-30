package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jussmor/blog/config"
	"github.com/jussmor/blog/internal/enums"
	"github.com/jussmor/blog/internal/web"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.GetEnv("JWT_SECRET")),
		ErrorHandler: jwtError,
		AuthScheme:   "Bearer",
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	fmt.Println(err)
	if err.Error() == "Missing or malformed JWT" {
		return web.JsonResponse(ctx, http.StatusForbidden, enums.ERROR_UNAUTHORIZED, nil)
	}
	return web.JsonResponse(ctx, http.StatusUnauthorized, enums.ERROR_JWT_EXPIRED, nil)
}
