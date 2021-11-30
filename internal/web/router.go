package web

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/jussmor/blog/infrastructures/db"
)

type RouterStruct struct {
	Web         *fiber.App
	PostgresDB   db.PostgresDB
}
