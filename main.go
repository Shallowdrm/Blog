package main

import (
	//"blog/email"
	"blog/api"
	"blog/database"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

func main() {

	database.Init()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	api.UserRouterInit(r)

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "loginPage.html", gin.H{})
	})

	r.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(200, "registerPage.html", gin.H{})
	})

	r.Run(":8080")
}

func Wmain() {

	//GET 用于获取信息，是无副作用的，是幂等的，且可缓存
	//POST 用于修改服务器上的数据，有副作用，非幂等，不可缓存
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/blog", func(ctx *gin.Context) {
		ctx.HTML(200, "default/index.html", gin.H{
			"time": fmt.Sprintf("%v", time.Now().Unix()),
		})
	})

	r.GET("/adduser", func(ctx *gin.Context) {
		ctx.HTML(200, "add_user.html", gin.H{})
	})

	r.POST("/doadduser", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		fmt.Println(username, password)
	})

	//email.Send() //send改成大写开头，go约定大写开头的标识符才能被import用
	//建议引用的函数大写开头
	log.Printf("Successfully")
	r.Run(":8080")
}
