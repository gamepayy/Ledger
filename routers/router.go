package routers

import (
	v1 "gamepayy_ledger/routers/api/v1"

	"github.com/gin-gonic/gin"

	docs "gamepayy_ledger/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := router.Group("/api/v1")
	{

		// @Summary 	Gets an account's data
		// @Produce  	json
		// @Param 		account query string true "Account"
		// @Success 	200 {object} object
		// @Failure 	400 {object} object
		// @Router 		/user [get]
		apiv1.GET("/user", v1.GetLedger)
		apiv1.GET("/token", v1.GetToken)

		apiv1.POST("/user/new", v1.NewLedger)
		apiv1.POST("/token/new", v1.NewToken)

		apiv1.DELETE("/user/delete", v1.DeleteLedger)
		apiv1.DELETE("/token/delete", v1.DeleteToken)

		apiv1.PUT("/user/transfer", v1.TransferLedger)
		apiv1.PUT("/user/withdraw", v1.WithdrawLedger)
		apiv1.PUT("/user/deposit", v1.DepositLedger)
		apiv1.PUT("/token/update", v1.UpdateToken)
	}

	router.Run("localhost:8080")
	return router
}
