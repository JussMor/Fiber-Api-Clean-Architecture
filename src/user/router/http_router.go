package router

import (
	"log"

	middlewares "github.com/jussmor/blog/internal/middlewares"
	"github.com/jussmor/blog/src/user/handlers"
	"github.com/jussmor/blog/src/user/repositories"
	"github.com/jussmor/blog/src/user/services"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Users Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	userPostgresDBRepo := repositories.NewUserRepostiory(r.PostgresDB)
	userService := services.NewUserService(userPostgresDBRepo, r.jwtAuth)
	authHandlers := handlers.NewHttpHandler(userService)
	userHandlers := handlers.NewUserHttpHandler(userService)

	r.Web.Post("/api/auth/login", authHandlers.Login)
	r.Web.Post("/api/auth/signup", userHandlers.CreateUser)
	r.Web.Get("/api/user/:id", userHandlers.GetUser)
	r.Web.Delete("/api/user", middlewares.Protected(), userHandlers.DeleteUser)
	r.Web.Put("/api/user", middlewares.Protected(), userHandlers.UpdateUser)
}
