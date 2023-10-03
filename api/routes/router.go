package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from golang vercel")
	})
}
