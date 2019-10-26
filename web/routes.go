package web

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// StartServer start the web server
func StartServer() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.Any("/graphql", graphQL)
		api.GET("/players", players)
		api.GET("/player_datas", playerDatas)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
