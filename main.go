package main

import (
	"net/http"

	"github.com/HamzaBenyazid/minitializr/service"
	"github.com/HamzaBenyazid/minitializr/types"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// health test
	r.GET("/heath", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.POST("/generate", func(c *gin.Context) {

		var miConfig types.MIConfig

		if err := c.Bind(&miConfig); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error(), "error_type": "Bind failure"})
			return
		}
		zipFile, err := service.Generate(&miConfig)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error(), "error_type": "Generate failure"})
			return
		}
		c.Header("Content-Disposition", "attachment; filename="+miConfig.Metadata["name"] + ".zip")
		c.Data(http.StatusOK, "application/zip", zipFile.Bytes())
	})

	return r
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
