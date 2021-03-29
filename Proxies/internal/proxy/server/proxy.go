package server

import (
	"Go-Gin/Proxies/internal/proxy/biz"
	"Go-Gin/Proxies/internal/proxy/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
@Time    : 2021/3/29 09:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : proxy.go
@Software: GoLand
*/
var business = biz.MakeProxyService(&data.ProxiesDaoImpl{})

// create handlerFunc to solve business
func ProxiesListHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 继承数据库操作对象
		c.JSON(http.StatusOK, gin.H{"id": 1, "ip": "192.168.31.1", "port": 8080})
	}
}

// other return reload gin.HandlerFunc
func ProxiesListHandler(c *gin.Context) {
	// 此处测试查询参数的验证器，但是没有使用数据库的查询
	search := biz.ProxySearchDTO{}
	err := c.Bind(&search)
	fmt.Println("参数绑定校验: ", err)

	if d, err := business.GetAllProxiesInfo(); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"data": nil, "count": 0, "msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"data": d, "count": len(d), "msg": "OK"})
	}
}

// create proxy to save database
func CreateProxyHandler(c *gin.Context) {
	// 将对象映射为DTO对象，传递进行处理
	dto := biz.ProxyDTO{}
	err := c.BindJSON(&dto)
	fmt.Println("接受到的数据:", dto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	err = business.AddOneProxy(&dto)
	var msg string = http.StatusText(http.StatusOK)
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

// delete proxy from database
func DeleteProxyHandler(c *gin.Context) {
	id := c.Param("id")
	proxyId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "msg": err.Error()})
		return
	}
	d := business.DeleteOneProxyById(proxyId)
	c.JSON(http.StatusOK, gin.H{"data": d, "msg": "OK"})
}

// search a proxy by id
func GetProxyByIdHandler(c *gin.Context) {
	id := c.Param("id")
	proxyId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "msg": err.Error()})
		return
	}
	d := business.GetProxyById(proxyId)
	c.JSON(http.StatusOK, gin.H{"data": d, "msg": "OK"})
}