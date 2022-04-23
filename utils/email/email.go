package email

import (
	"admin-cli/global"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

// Email Email发送方法
func Email(To, subject string, body string) error {
	to := strings.Split(To, ",")
	return send(to, subject, body)
}

// ErrorToEmail 给email中间件错误发送邮件到指定邮箱
func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.Config.Email.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

// EmailTest Email测试方法
func EmailTest(subject string, body string) error {
	to := []string{global.Config.Email.From}
	return send(to, subject, body)
}

//Email发送方法
func send(to []string, subject string, body string) error {
	from := global.Config.Email.From
	nickname := global.Config.Email.Nickname
	secret := global.Config.Email.Secret
	host := global.Config.Email.Host
	port := global.Config.Email.Port
	isSSL := global.Config.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
