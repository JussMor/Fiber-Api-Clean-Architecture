package main 

import(
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2" 
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cors"
	db "github.com/jussmor/blog/infrastructures/db"
	"github.com/jussmor/blog/internal/routes"
	"github.com/jussmor/blog/internal/web"

)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			if err != nil {
				return web.JsonResponse(ctx, code, "Error del servidor", nil)
			}

			return nil
		},
	})

	app.Use(recover.New())

	app.Use(limiter.New(limiter.Config{
		Next:         func(c *fiber.Ctx) bool { return false },
		Max:          60,
		Expiration:   10 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string { return c.Get("x-forwarded-for") },
		LimitReached: func(c *fiber.Ctx) error {
			return web.JsonResponse(c, http.StatusTooManyRequests, "Estás accediendo demasiado rápido.", nil)
		},
	}))

	app.Use(cors.New())

	postgresDB := db.ConnectDB()

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       app,
			PostgresDB: postgresDB,
		},
	}

	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()

	app.Listen(":3000")

}