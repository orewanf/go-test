package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

func main() {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	initializeRoutes(router)

	// Run the program
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
