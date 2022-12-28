package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/sai1024/bit_backend/docs"
	"github.com/sai1024/bit_backend/middleware/jwt"
	"github.com/sai1024/bit_backend/pkg/setting"
	"github.com/sai1024/bit_backend/routers/api"
	v1 "github.com/sai1024/bit_backend/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/manage/auth", api.Auth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/redeem_codes", v1.GetRedeemCodes)
		apiv1.POST("/redeem_codes", v1.AddRedeemCode)
		apiv1.DELETE("/redeem_codes/:id", v1.DeleteRedeemCode)
	}

	return r
}
