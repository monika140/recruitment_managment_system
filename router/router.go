package router

import (
	"Go_authentication/controllers"
	"Go_authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	authRoutes := r.Group("/auth") //Creates a group of auth routes.
	{
		authRoutes.POST("/register", controllers.Register) //Route for user registration.
		authRoutes.POST("/login", controllers.Login)       //Route for user login.
		authRoutes.GET("/:id", controllers.GetUser)
		authRoutes.PUT("/:id", controllers.UpdateUser)
		authRoutes.DELETE("/:id", controllers.DeleteUser)
	}

    
	applicant := r.Group("/applicant")
	{
		applicant.GET("/jobs", controllers.ListJobs)
		applicant.GET("/jobs/apply", controllers.ApplyToJob)
	}
    
	admin := r.Group("/admin")
	{
		admin.Use(middlewares.AuthMiddleware()) // Apply the middleware
		admin.POST("/job", controllers.CreateJob)
	}

	return r
}
