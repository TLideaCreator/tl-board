package router

import (
	"github.com/gin-gonic/gin"
)

func RouteRegister() {
	router := gin.New()

	router.Group("")

	_ = router.Run(":10010")
}
