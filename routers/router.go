package routers

import (
	v1 "gamepayy_ledger/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	apiv1 := router.Group("/api/v1")
	{

		apiv1.GET("/user", v1.GetLedger)

		apiv1.POST("/user/new", v1.NewLedger)
		apiv1.POST("/new_token", v1.NewToken)

		apiv1.DELETE("/user/delete", v1.DeleteLedger)
		apiv1.DELETE("/delete_token", v1.DeleteToken)

		apiv1.PUT("/edit_token", v1.EditToken)
		apiv1.PUT("/user/transfer", v1.TransferLedger)
		apiv1.PUT("/user/deposit", v1.DepositLedger)
		apiv1.PUT("/user/withdraw", v1.WithdrawLedger)
	}

	//router.Run("localhost:8080")
	return router
}
