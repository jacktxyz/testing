package srv

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			fmt.Println("alive")
		}
	}()

	r.GET("/v1/binding", BindingV1)

	r.POST("/v1/hooks", Hooks)

	r.Run(":9990")
}

func Hooks(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, err)
		return
	}

	fmt.Println(string(data))
	c.JSON(200, "ok")
}

func BindingV1(c *gin.Context) {

	q := QueryWrapper{}
	if err := c.ShouldBindJSON(&q); err != nil {
		fmt.Println("error message: ", err)
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	c.JSON(http.StatusOK, "ok v1 hello dev")
}

// overwrite
type QueryWrapper struct {
	Query
	Name      string `json:"name"`
	ExpiredAt string `json:"expiredAt"`
}

type Query struct {
	ExpiredAt *time.Time `json:"expiredAt" form:"expiredAt"`
	Name      string     `json:"name"`
	Addr      string     `json:"addr"`
}
