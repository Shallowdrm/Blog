package email

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func Send(receiver string) {
	em := email.NewEmail()

	sender := "2750334729@qq.com"
	//receiver := "651376825@qq.com"

	em.From = fmt.Sprintf("Bot <%s>", sender)
	//设置发送方邮箱	em.From = "Test <651376825@qq.com>"

	em.To = []string{receiver}
	//设置接收方邮箱	em.To = []string{"2750334729@qq.com"}

	em.Subject = "测试邮件-验证码测试"
	//设置主题

	//html := ``
	//em.HTML = []byte(html)
	//发送HTML文件

	rand.Seed(time.Now().Unix())
	data := fmt.Sprintf("%06v", rand.Intn(999999))

	em.Text = []byte(data)
	//设置邮件发送的内容(纯文本)

	err := em.Send(
		"smtp.qq.com:25", //参数1是SMTP服务器的地址，参数2为验证信息。
		smtp.PlainAuth(
			"",
			sender,             //"651376825@qq.com",
			"upouuohfoirodhfb", //"unqyxdxfpynxbbac",
			"smtp.qq.com"),
	)

	if err != nil {
		log.Fatal(err)
		//日志信息
	}

	log.Println("send successfully")
	//日志信息
}
