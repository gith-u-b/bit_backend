package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/sai1024/bit_backend/pkg/setting"
	"github.com/sai1024/bit_backend/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	return r
}
