package router

import (
	auth "github.com/jussmor/blog/internal/utils"
	"github.com/jussmor/blog/internal/web"
)

type RouterStruct struct {
	web.RouterStruct
	jwtAuth auth.JwtTokenInterface
}