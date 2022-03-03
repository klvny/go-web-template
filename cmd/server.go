package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// router.Use(cors.Default())

	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, "hello")
		})
	}

	fmt.Println("Port: ", AppCfg.App.Port)
	router.Run(":" + AppCfg.App.Port)
}
