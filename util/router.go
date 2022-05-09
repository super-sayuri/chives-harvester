package util

import "github.com/gin-gonic/gin"

type ReturnApi struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ResponseError(c *gin.Context, code int, err error) {
	returns := &ReturnApi{
		Code:    code,
		Data:    err.Error(),
		Message: err.Error(),
	}
	c.JSON(code, returns)
}

func NormalResponse(c *gin.Context, data interface{}) {
	returns := &ReturnApi{
		Code:    200,
		Data:    data,
		Message: "SUCCESS",
	}
	c.JSON(200, returns)
}
