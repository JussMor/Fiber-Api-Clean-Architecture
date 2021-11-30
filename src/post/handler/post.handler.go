package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jussmor/blog/internal/enums"
	"github.com/jussmor/blog/internal/utils"
	"github.com/jussmor/blog/internal/web"
	"github.com/jussmor/blog/src/post/dto"
	"github.com/jussmor/blog/src/post/services"
)

type PostHandlerInterface interface {
	GetPost(ctx *fiber.Ctx) error
	GetAllPost(ctx *fiber.Ctx) error
	CreatePost(ctx *fiber.Ctx) error
	UpdatePost(ctx *fiber.Ctx) error
	DeletePost(ctx *fiber.Ctx) error
}

type PostHandler struct {
	postService services.PostServiceInterface
}

func NewUserHttpHandler(
	postService services.PostServiceInterface,
) PostHandlerInterface {
	return &PostHandler{
		postService: postService,
	}
}
func (services *PostHandler) GetAllPost(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page"))

	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))

	if err != nil {
		pageSize = 10
	}

	postList := services.postService.FindByAll(page, pageSize)

	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, postList)
}

func (services *PostHandler) GetPost(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	post := services.postService.FindById(uint(id))

	if post.UserID == 0 {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.MSG_POST_NOT_FOUND, nil)
	}
	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, post)
}

func (services *PostHandler) CreatePost(ctx *fiber.Ctx) error {

	newPost := new(dto.Post)

	if err := ctx.BodyParser(&newPost); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)

	}

	errors := utils.Validate(newPost)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	res, err := services.postService.Save(*newPost)

	if err != nil {
		web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_SERVER, nil)
	}
	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, res)
}

func (services *PostHandler) UpdatePost(ctx *fiber.Ctx) error {
	postUpdate := new(dto.PostUpdate)

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_ID_NOT_FOUND, nil)
	}

	if err := ctx.BodyParser(&postUpdate); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	errors := utils.Validate(postUpdate)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	category, err := services.postService.Update(uint(id), *postUpdate)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_UPDATE, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, category)
}

func (services *PostHandler) DeletePost(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_DELETE, nil)
	}

	err = services.postService.Delete(uint(id))

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_DELETE, nil)
	} else {
		return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, nil)
	}
}
