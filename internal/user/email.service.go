package user

import (
	"demo01/pkg/logger/logrusx"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func (a *service) Send(ctx *gin.Context) {

	m := gomail.NewMessage()
	m.SetHeader("From", "xxx@qq.com")
	m.SetHeader("To", "xxx@gmail.com", "xx@foxmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	d := gomail.NewDialer("smtp.qq.com", 465, "用户名", "密码")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		logrusx.Info(ctx, err)
	}
}
