package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type BaseController struct{}

type result struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (con BaseController) success(c *gin.Context, data any) {
	fmt.Println(data)
	res := result{Code: 200, Data: data}
	fmt.Println(res)
	c.JSON(200, res)
}

func (con BaseController) fail(c *gin.Context, msg string) {
	res := result{Code: 500, Msg: msg}
	c.JSON(200, res)
}
