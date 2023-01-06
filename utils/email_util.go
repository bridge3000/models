package utils

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"regexp"
	"strings"
)

type EmailUtil struct {
	ServerHost   string
	ServerPort   int
	FromEmail    string
	FromPassword string
	Subject      string
	Body         string
	Toers        string
	CCers        string
}

// SendEmail body支持html格式字符串
func (this *EmailUtil) Send() error {
	toers := []string{}

	m := gomail.NewMessage()

	if len(this.Toers) == 0 {
		return fmt.Errorf("Empty toers")
	}

	for _, tmp := range strings.Split(this.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toers...)

	//抄送列表
	if len(this.CCers) != 0 {
		for _, tmp := range strings.Split(this.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", toers...)
	}

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	//	m.SetAddressHeader("From", "s-game")
	m.SetAddressHeader("From", this.FromEmail, "s-game")

	// 主题
	m.SetHeader("Subject", this.Subject)

	//	m.SetHeader("Reply-To", "qiaoliang@s-game.com.cn")

	// 正文
	m.SetBody("text/html", this.Body)

	d := gomail.NewPlainDialer(this.ServerHost, this.ServerPort, this.FromEmail, this.FromPassword)

	// 发送
	err := d.DialAndSend(m)
	return err
}

// 识别电子邮箱
func (this *EmailUtil) IsEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`

	emailLen := len(email)
	if emailLen > 10 {
		domain := email[emailLen-9:]

		if domain == "gmail.com" {
			//			pattern = `^[a-z0-9.]+@gmail.com`
			pattern = `^[A-Za-z0-9.]+@gmail.com`
		}
	}

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func (this *EmailUtil) SendWarning(subject string, body string) {
	//	this.Toers = "9455856@qq.com,13811557683@139.com"
	this.Toers = "qiaoliang@s-game.com.cn"
	this.Subject = subject
	this.Body = body
	this.Send()
}
