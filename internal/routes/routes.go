package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jussmor/blog/internal/web"
	categoryRouter "github.com/jussmor/blog/src/category/router"
	postRouter "github.com/jussmor/blog/src/post/router"
	tagRouter "github.com/jussmor/blog/src/tag/router"
	userRouter "github.com/jussmor/blog/src/user/router"
)

type RouterStruct struct {
	web.RouterStruct
}

func NewHttpRoute(r RouterStruct) RouterStruct {
	log.Println("Loading the HTTP Router")

	return r
}

func (c *RouterStruct) GetRoutes() {
	c.Web.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello this is my first route in go fiber"))
	})

	webRouterConfig := web.RouterStruct{
		Web:       c.Web,
		PostgresDB : c.PostgresDB ,
	}
	// registering route from another modules
	userRouterStruct := userRouter.RouterStruct{
		RouterStruct: webRouterConfig,
	}
	userRouter := userRouter.NewHttpRoute(userRouterStruct)
	userRouter.GetRoute()

	
	categoryRouterStruct := categoryRouter.RouterStruct{
		RouterStruct: webRouterConfig,
	}
	categoryRouter := categoryRouter.NewHttpRoute(categoryRouterStruct)
	categoryRouter.GetRoute()

	tagRouterStruct := tagRouter.RouterStruct{
		RouterStruct: webRouterConfig,
	}
	tagRouter := tagRouter.NewHttpRoute(tagRouterStruct)
	tagRouter.GetRoute()

	postRouterStruct := postRouter.RouterStruct{
		RouterStruct: webRouterConfig,
	}
	postRouter := postRouter.NewHttpRoute(postRouterStruct)
	postRouter.GetRoute()

	// handling 404 error
	c.Web.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
}
