package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/v1/binding", BindingV1)

	v2Group := r.Group("/v2/binding").Use(AuthV2)
	v2Group.GET("/", BindingV2)

	r.Run(":9990")
}

func BindingV1(c *gin.Context) {

	c.JSON(http.StatusOK, "ok v1 01")
}

func BindingV2(c *gin.Context) {

	c.JSON(http.StatusOK, "ok v2")
}

func AuthV2(c *gin.Context) {
	fmt.Println("v2 authorized")
}
