package router

import "github.com/gin-gonic/gin"

type Demo struct{}

func (demo Demo) Route(rg *gin.RouterGroup) {
	rg.GET("/get-user-info", demo.getUserInfo)
}

func (Demo) getUserInfo(ctx *gin.Context) {

}
