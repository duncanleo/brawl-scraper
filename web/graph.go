package web

import (
	"github.com/duncanleo/brawl-scraper/graph"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

var (
	schema         = graph.Schema()
	graphQLHandler = handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: gin.IsDebugging(),
	})
)

func graphQL(c *gin.Context) {
	graphQLHandler.ServeHTTP(c.Writer, c.Request)
}
