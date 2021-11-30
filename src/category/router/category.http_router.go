
package router
import (

	"log"

	middlewares "github.com/jussmor/blog/internal/middlewares"
	"github.com/jussmor/blog/src/category/handler"
	"github.com/jussmor/blog/src/category/repositories"
	"github.com/jussmor/blog/src/category/services"

)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Category Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	categoryPostgresDBRepo := repositories.NewUserRepostiory(r.PostgresDB)
	categoryService := services.NewUserService(categoryPostgresDBRepo)
	categoryHandler := handler.NewUserHttpHandler(categoryService)

	r.Web.Get("/api/category", categoryHandler.GetAllCategory)
	r.Web.Get("/api/category/:id", categoryHandler.GetCategory)
	r.Web.Post("/api/category", middlewares.Protected(), categoryHandler.CreateCategory)
	r.Web.Put("/api/category/:id", middlewares.Protected(), categoryHandler.UpdateCategory)
	r.Web.Delete("/api/category/:id", middlewares.Protected(), categoryHandler.DeleteCategory)
}
