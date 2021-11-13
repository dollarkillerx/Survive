package utils

import (
	"fmt"

	"github.com/dollarkillerx/survive/internel/conf"
	"github.com/dollarkillerx/survive/pkg/models"
	"gopkg.in/gomail.v2"
)

type Email struct {
	conf conf.EmailConfig
}

func NewEmail(conf conf.EmailConfig) *Email {
	return &Email{
		conf: conf,
	}
}

func (e *Email) SendEmail(node models.Node) error {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", fmt.Sprintf("Survive <%s>", e.conf.Email))
	//接收人
	m.SetHeader("To", e.conf.Email)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "Survive 服务器告警")
	//内容
	m.SetBody("text/html", fmt.Sprintf(emailTemp, node.IP, node.Delay, node.Name, node.Token, node.LastPingTime, node.CPU, node.CPUUsedPercent, node.MemUsedPercent))
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(e.conf.SMTPHost, e.conf.SMTPort, e.conf.Email, e.conf.Password)

	return d.DialAndSend(m)
}

var emailTemp = `
<h1>Survive 服务器告警</h1>

Server IP:  %s,  延迟: %d,  Name: %s ,  Token: %s , LastPingTime: %s, CPU: %s , CPUUse: %s  , MemUse: %s, 
`
