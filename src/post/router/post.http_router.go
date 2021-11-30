package router

import (
	"log"

	middlewares "github.com/jussmor/blog/internal/middlewares"
	"github.com/jussmor/blog/src/post/handler"
	"github.com/jussmor/blog/src/post/repositories"
	"github.com/jussmor/blog/src/post/services"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Tag Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	postPostgresDB := repositories.NewUserRepostiory(r.PostgresDB)
	postService := services.NewUserService(postPostgresDB)
	postHandler := handler.NewUserHttpHandler(postService)

	r.Web.Get("/api/post", postHandler.GetAllPost)
	r.Web.Get("/api/post/:id", postHandler.GetPost)
	r.Web.Post("/api/post", middlewares.Protected(), postHandler.CreatePost)
	r.Web.Put("/api/post/:id", middlewares.Protected(), postHandler.UpdatePost)
	r.Web.Delete("/api/post/:id", middlewares.Protected(), postHandler.DeletePost)
}
