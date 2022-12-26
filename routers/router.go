package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/sai1024/bit_backend/docs"
	"github.com/sai1024/bit_backend/pkg/setting"
	"github.com/sai1024/bit_backend/routers/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/auth", api.GetAuth)

	return r
}
