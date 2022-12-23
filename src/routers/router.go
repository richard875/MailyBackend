package routers

import (
	"maily/go-backend/src/controllers"
	"maily/go-backend/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// Suffix api
	public := router.Group("/api")

	// GET
	public.GET("/beep/:trackingId", controllers.Beep)    // Logging
	public.GET("/ip-test", controllers.IpAddress)        // Dev Test
	public.GET("/browser-test", controllers.BrowserTest) // Dev Test

	// POST
	public.POST("/register", controllers.Register) // Auth
	public.POST("/login", controllers.Login)       // Auth
	public.POST("/generate", controllers.Generate) // Generate tracking number

	// Suffix admin
	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())

	// GET
	protected.GET("/user", controllers.CurrentUser)
}
