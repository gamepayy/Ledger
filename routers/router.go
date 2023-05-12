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

	user := router.Group("/api/v1/user")
	{
		user.GET("", v1.GetLedger)

		user.POST("/new", v1.NewLedger)

		user.DELETE("/delete", v1.DeleteLedger)

		user.PUT("/transfer", v1.TransferLedger)

		user.PUT("/withdraw", v1.WithdrawLedger)

		user.PUT("/deposit", v1.DepositLedger)

	}

	token := router.Group("/api/v1/token")
	{
		token.GET("", v1.GetToken)

		token.POST("/new", v1.NewToken)

		token.DELETE("/delete", v1.DeleteToken)

		token.PUT("/update", v1.UpdateToken)

	}

	withdraws := router.Group("/api/v1/withdraws")
	{
		withdraws.GET("", v1.GetWithdraws)

		withdraws.POST("/new", v1.InsertWithdraw)

		withdraws.DELETE("/delete", v1.DeleteWithdraw)

		withdraws.DELETE("/clean", v1.DeleteProcessedWithdraws)

		withdraws.PUT("/process", v1.ProcessWithdraw)
	}

	router.Run("localhost:8080")
	return router
}
