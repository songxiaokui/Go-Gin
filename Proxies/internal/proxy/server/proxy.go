package server

import (
	"Go-Gin/Proxies/internal/proxy/biz"
	"Go-Gin/Proxies/internal/proxy/data"
	myerror "Go-Gin/Proxies/internal/proxy/pkg/error"
	myrsp "Go-Gin/Proxies/internal/proxy/pkg/responseJson"
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
	//err := c.Bind(&search)
	//fmt.Println("参数绑定校验: ", err)
	//
	//if d, err := business.GetAllProxiesInfo(); err != nil {
	//	c.IndentedJSON(http.StatusOK, gin.H{"data": nil, "count": 0, "msg": err.Error()})
	//} else {
	//	c.IndentedJSON(http.StatusOK, gin.H{"data": d, "count": len(d), "msg": "OK"})
	//}
	// 使用自定义错误处理
	myerror.MakeMultiError(c.Bind(&search)).Unwrap()
	dataList := myerror.MakeMultiError(business.GetAllProxiesInfo()).Unwrap()
	// 能走到这里，说明肯定成功，直接将结果包装返回,使用自定义json返回
	myrsp.R(c)("200001", "OK", dataList)(myrsp.OK)

}

// create proxy to save database
func CreateProxyHandler(c *gin.Context) {
	// 将对象映射为DTO对象，传递进行处理
	//dto := biz.ProxyDTO{}
	//err := c.BindJSON(&dto)
	//fmt.Println("接受到的数据:", dto)
	//if err != nil {
	//	// c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	//	// 使用middleware封装的捕获异常操作
	//	panic(err.Error())
	//	return
	//}
	//err = business.AddOneProxy(&dto)
	//if err != nil {
	//	panic(err.Error())
	//}
	//c.JSON(http.StatusOK, gin.H{"msg": http.StatusText(http.StatusOK)})

	// 使用自定义error包装单错误返回对象
	dto := biz.ProxyDTO{}
	myerror.MakeSignalError(c.BindJSON(&dto)).Unwrap()
	myerror.MakeSignalError(business.AddOneProxy(&dto)).Unwrap()
	// 使用自定义json封装处理返回响应
	// c.JSON(http.StatusOK, gin.H{"msg": http.StatusText(http.StatusOK)})
	myrsp.R(c)("2000001", "OK", nil)(myrsp.OK)
}

// delete proxy from database
func DeleteProxyHandler(c *gin.Context) {
	//id := c.Param("id")
	//proxyId, err := strconv.Atoi(id)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"data": nil, "msg": err.Error()})
	//	return
	//}
	//d := business.DeleteOneProxyById(proxyId)
	//c.JSON(http.StatusOK, gin.H{"data": d, "msg": "OK"})

	// 简化后代码
	id := c.Param("id")
	proxyId := myerror.MakeMultiError(strconv.Atoi(id)).Unwrap()
	myerror.MakeMultiError(business.DeleteOneProxyById(proxyId.(int)))
	// c.JSON(http.StatusOK, gin.H{"msg": "OK"})
	myrsp.R(c)("2000003", "OK", nil)(myrsp.OK)

}

// search a proxy by id
func GetProxyByIdHandler(c *gin.Context) {
	//id := c.Param("id")
	//proxyId, err := strconv.Atoi(id)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"data": nil, "msg": err.Error()})
	//	return
	//}
	//d := business.GetProxyById(proxyId)
	//c.JSON(http.StatusOK, gin.H{"data": d, "msg": "OK"})

	//// 使用简化了错误处理和响应处理代码
	id := c.Param("id")
	proxyId := myerror.MakeMultiError(strconv.Atoi(id)).Unwrap()
	myrsp.R(c)("2000001", "OK", business.GetProxyById(proxyId.(int)))(myrsp.OK)

}
