package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

////////////////////

// HelloController handle the Hello Controller
type HelloController struct {
}

// Setup configures the routes for this controller
func (h *HelloController) Setup(router *gin.Engine) {
	router.GET("/hello", h.handleHello)
}

func (h *HelloController) handleHello(c *gin.Context) {
	c.String(200, "Hello World")
}

////////////////////

func handleHelloName(c *gin.Context) {
	name := c.Param("name")
	c.String(200, fmt.Sprintf("Hello World %s", name))
}

func main() {
	fmt.Println("Hello")

	router := gin.Default()

	hc := &HelloController{}
	hc.Setup(router)

	router.GET("/hello/:name", handleHelloName)

	router.Run()
}
