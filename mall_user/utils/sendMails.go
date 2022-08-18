package utils

import (
	"fmt"
	beegoUtils "github.com/astaxie/beego/utils"
	"strings"
)

func SendEmail(toEmail, message string) error {
	username := "903086461@qq.com"
	password := "alqxjwramgulbaie"
	host := "smtp.qq.com"
	port := "587"

	emailConfig := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%s}`, username, password, host, port)
	fmt.Println("emailConfig", emailConfig)
	emailConn := beegoUtils.NewEMail(emailConfig) // beego下的
	emailConn.From = strings.TrimSpace(username)
	emailConn.To = []string{strings.TrimSpace(toEmail)}
	emailConn.Subject = "mall邮箱验证码"
	//注意这里我们发送给用户的是激活请求地址
	emailConn.Text = message
	err := emailConn.Send()
	if err != nil {
		fmt.Println("send email err", err)
		return err
	}
	return nil

}
