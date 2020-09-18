//+build windows

package follower

import (
	"follower/controller"
	"follower/util/dao"
	"github.com/gin-gonic/gin"
	"log"
)

type Config struct {
	Addr string
}

func Run(conf Config) {
	dao.Startup()
	router := gin.New()
	controller.Router(&router.RouterGroup)
	err := router.Run(conf.Addr)
	if err != nil {
		log.Fatal("follower run err!")
	}
}
