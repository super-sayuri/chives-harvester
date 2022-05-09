package router

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"sayuri_crypto_bot/util"
)

func ImplementRouter(g *gin.RouterGroup) {
	baseUri := GetUri(ApiRouter)
	uriWithId := baseUri + "/:id"
	g.GET(baseUri, func(c *gin.Context) {
		items, err := listALL()
		if err != nil {
			util.ResponseError(c, 500, err)
			return
		}
		util.NormalResponse(c, items)
	})
	g.GET(uriWithId, func(c *gin.Context) {
		id := c.Param("id")
		items, err := listOne(id)
		if err != nil {
			util.ResponseError(c, 500, err)
			return
		}
		if items == nil {
			util.ResponseError(c, 404, errors.New("not found"))
			return
		}
		util.NormalResponse(c, items)
	})
	g.POST(baseUri, func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			util.ResponseError(c, 400, err)
			return
		}
		obj := &CRouter{
			Name:  "",
			Value: "",
		}
		err = json.Unmarshal(data, obj)
		if err != nil {
			util.ResponseError(c, 400, err)
			return
		}
		err = obj.SelfCheck(true)
		if err != nil {
			util.ResponseError(c, 400, err)
			return
		}
		err = insert(obj)
		if err != nil {
			util.ResponseError(c, 500, err)
			return
		}
		util.NormalResponse(c, "success")
	})
	g.PUT(uriWithId, func(c *gin.Context) {
		id := c.Param("id")
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			util.ResponseError(c, 400, err)
			return
		}
		obj := &CRouter{
			Name:  "",
			Value: "",
		}
		err = json.Unmarshal(data, obj)
		if err != nil {
			util.ResponseError(c, 400, err)
			return
		}
		err = obj.SelfCheck(false)
		if err != nil {
			util.ResponseError(c, 400, err)
			return
		}
		old, err := listOne(id)
		if err != nil {
			util.ResponseError(c, 400, err)
			return
		}
		if old == nil {
			util.ResponseError(c, 404, errors.New("not found"))
			return
		}
		old.Name = ""
		old.Value = obj.Value
		err = update(id, old)
		if err != nil {
			util.ResponseError(c, 500, err)
			return
		}
		util.NormalResponse(c, "success")
	})
	g.DELETE(uriWithId, func(c *gin.Context) {
		id := c.Param("id")
		err := deleteOne(id)
		if err != nil {
			util.ResponseError(c, 500, err)
			return
		}
		util.NormalResponse(c, "success")
	})
}
