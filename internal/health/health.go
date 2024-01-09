package health

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartHealthCheckServer(port int) {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	log.Printf("Starting health check server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
