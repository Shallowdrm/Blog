package api

import (
	"blog/database"
	"blog/email"
	"blog/model"
	"log"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	//创建User路由组
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", Register)
		userRouter.POST("/login", Login)
		userRouter.POST("/email", Email)
	}

}

//注册函数
func Register(c *gin.Context) {
	registerInfo := &model.RegisterInfo{}
	err := c.ShouldBind(registerInfo)
	if err != nil {
		//log.Fatal(err)
		log.Println(err)
	}

	ok := database.RS.CheckByUsername(registerInfo)
	//判断注册的用户名是否重复

	if ok {
		c.JSON(200, gin.H{
			"status": false,
			"data":   "用户名已存在",
		})
	} else {

		err = database.RS.CreatNewUser(registerInfo)
		if err != nil {
			//log.Fatal(err)
			log.Println(err)
		}

		log.Println("register successfully\n", registerInfo)
		c.JSON(200, gin.H{
			"status": true,
			"data":   "注册成功",
		})
	}

}

//登录函数
func Login(c *gin.Context) {
	loginInfo := &model.LoginInfo{}
	err := c.ShouldBind(loginInfo)
	if err != nil {
		log.Fatal(err)
	}

	user, ok := database.RS.FindUserById(loginInfo.Username)

	if !ok {
		c.JSON(200, gin.H{
			"status": false,
			"data":   "用户名不存在",
		})
	} else if user.Password != loginInfo.Password {
		c.JSON(200, gin.H{
			"status": false,
			"data":   "密码错误",
		})
	} else {

		log.Println("login successfully\n", loginInfo)

		c.JSON(200, gin.H{
			"status": true,
			"data":   "登陆成功",
		})
	}

}

//发送邮件函数
func Email(c *gin.Context) {
	receiver := c.PostForm("email")
	email.Send(receiver)
	c.JSON(200, gin.H{
		"status": "true",
		"data":   "验证码发送成功",
	})
}
