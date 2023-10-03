package api

import (
	"net/http"

	"vercelgin/api/routes"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()
	r := app.Group("/api")
	routes.Router(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
