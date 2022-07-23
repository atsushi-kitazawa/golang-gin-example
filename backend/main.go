package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PasswordChanger interface {
}

type Data struct {
	Id   int
	Name string
}

func main() {
	doMain()
}

func doMain() {
	r := gin.Default()

	// add route
	getExam1(r)
	getExam2(r)

	r.Run(":8080")
}

func getExam1(r *gin.Engine) {
	r.GET("/datas", func(c *gin.Context) {
		datas := make([]Data, 0)
		datas = append(datas, Data{Id: 1, Name: "foo"})
		datas = append(datas, Data{Id: 2, Name: "bar"})
		err := c.BindJSON(datas)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.String(http.StatusOK, fmt.Sprintf("%v", datas))
	})
}

func getExam2(r *gin.Engine) {
	r.GET("/datas/:id", func(c *gin.Context) {
		_ = c.Param("id")
		data := Data{Id: 1, Name: "foo"}
		err := c.BindJSON(data)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.String(http.StatusOK, fmt.Sprintf("%v", data))
	})
}
