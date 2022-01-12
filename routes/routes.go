package routes

import (
	"go-contact/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	// menandaka router
	r := gin.Default()

	// set db to gin context -> that in can acess globally i
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/contact", controllers.GetAllContact)
	r.POST("/contact", controllers.CreateContact)
	r.GET("/contact/:id", controllers.GetContactById)
	r.PATCH("/contact/:id", controllers.UpdateContact)
	r.DELETE("/contact/:id", controllers.DeleteContact)

	return r
}
