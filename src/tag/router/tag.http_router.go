package router

import (
	"log"

	middlewares "github.com/jussmor/blog/internal/middlewares"
	"github.com/jussmor/blog/src/tag/handler"
	"github.com/jussmor/blog/src/tag/repositories"
	"github.com/jussmor/blog/src/tag/services"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Tag Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	tagPostgresDB := repositories.NewUserRepostiory(r.PostgresDB)
	tagService := services.NewUserService(tagPostgresDB )
	tagHandler := handler.NewUserHttpHandler(tagService)

	r.Web.Get("/api/tag", tagHandler.GetTag)
	r.Web.Post("/api/tag", middlewares.Protected(), tagHandler.CreateTag)
	r.Web.Put("/api/tag/:id", middlewares.Protected(), tagHandler.UpdateTag)
	r.Web.Delete("/api/tag/:id", middlewares.Protected(), tagHandler.DeleteTag)
}
