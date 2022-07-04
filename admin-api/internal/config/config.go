package config

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host string
		Pass string
	}
	Mysql struct {
		Datasource string
	}
}

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}

// 邮箱验证码发送
func MailSendCode(fmail, tmail, subject, code, username, password, host, port string) error {
	e := email.NewEmail()
	//e.From = "<hzy15847445046@163.com>"
	//e.To = []string{mail}
	//e.Subject = "宏筑建通"
	//e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	//err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("",
	//	"hzy15847445046@163.com",
	//	"HTZMPBRBHVEEDLLP",
	//	"smtp.163.com"),
	//	&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	//if err != nil {
	//	return err
	//}
	e.From = fmail
	e.To = []string{tmail}
	e.Subject = subject
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	return e.SendWithTLS(host+":"+port,
		smtp.PlainAuth(
			"",
			username,
			password,
			host,
		),
		&tls.Config{InsecureSkipVerify: true, ServerName: host},
	)
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
