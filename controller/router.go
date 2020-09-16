package controller

import (
	"follower/controller/router"
	"github.com/gin-gonic/gin"
)

func Router(rg *gin.RouterGroup) {
	router.Demo{}.Route(rg.Group("/demo"))
}
