package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//使用默认中间件创建一个gin路由器
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/somePost", postting) //这个路由区分大小写
	//匹配地址参数
	r.GET("/user/:name", user)
	r.GET("/user/:name/*action", users)
	//获取get参数
	r.GET("/welcome", welcome)
	//获取post参数
	r.POST("/form", form)
	//文件上传
	r.POST("/upload", upload)

	//路由分组
	api := r.Group("/api")
	{
		api.POST("/login", login)
		api.POST("/register", register)
	}

	r.Run() //默认监听8080端口 也可以自己定义
}

func login(c *gin.Context) {
	username, _ := c.GetPostForm("username") //多返回一个key是否存在的结果
	password, _ := c.GetPostForm("password")
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"username": username,
		"password": password,
	})
}

func register(c *gin.Context) {
	username := c.DefaultPostForm("username", "uname")
	password := c.PostForm("password")
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"username": username,
		"password": password,
	})
}

//文件上传
func upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	uploadPath := "./upload/" + file.Filename
	//上传文件到制定路径
	err := c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		fmt.Println("upload err:", err)
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%s upload!", file.Filename))
}

//获取Get参数
func welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") //是c.Request.URL.Query.Get("lastname")的简写
	c.String(http.StatusOK, "hello %s %s", firstname, lastname)
}

//获取Post参数
func form(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous") //此方法可以设置默认值
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func user(c *gin.Context) {
	name := c.Param("name") //
	c.String(http.StatusOK, "Hello %s", name)
}

func users(c *gin.Context) { // /dog/wangwang dogis /wangwang
	name := c.Param("name")
	action := c.Param("action")
	message := name + "is " + action
	c.String(http.StatusOK, message)
}
func postting(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "post",
	})
}
