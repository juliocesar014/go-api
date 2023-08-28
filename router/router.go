package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Routers Default
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	//  Run Server
	router.Run(":8080")
}
