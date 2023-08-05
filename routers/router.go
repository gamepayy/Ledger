package routers

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"

	docs "gamepayy_ledger/docs"
	jwt "gamepayy_ledger/middleware/jwt"
	logging "gamepayy_ledger/middleware/logging"
	v1 "gamepayy_ledger/routers/api/v1"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func InitRouter() *gin.Engine {
	router := gin.New()

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 10,
	})
	rateLimitMiddleware := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	router.Use(rateLimitMiddleware)

	router.Use(logging.GinzapLogger())

	authMiddleware, err := jwt.CreateAuthMiddleware()
	if err != nil {
		panic("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		panic("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)
	auth := router.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", jwt.HelloHandler)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := router.Group("/api/v1/user")
	{
		user.GET("", v1.GetLedger)

		user.POST("/new", v1.NewLedger)

		user.DELETE("/delete", v1.DeleteLedger)

		user.PUT("/transfer", v1.TransferLedger)

		user.POST("/withdraw", v1.WithdrawLedger)

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

	return router
}
