package remindmetools
import (
    // "crypto/tls"
    "fmt"
    "net/smtp"
)

type mail struct {
    user   string
    passwd string
}

func check(err error) bool{
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

//初始化用户名和密码
func New(u string, p string) mail {
    temp := mail{user: u, passwd: p}
    return temp
}

//标题 文本 目标邮箱
func (m mail) Send(title string, text string, toId string) bool{
    auth := smtp.PlainAuth("", m.user, m.passwd, "smtp.gmail.com")
    // fmt.Printf("useraccount:%v\npassword:%v\ntoemail:%v\n",m.user,m.passwd,toId)
    // tlsconfig := &tls.Config{
    //     InsecureSkipVerify: true,
    //     ServerName:         "smtp.gmail.com",
    // }

    // conn, err := tls.Dial("tcp", "smtp.gmail.com:465", tlsconfig)
    // check(err)

    // client, err := smtp.NewClient(conn, "smtp.gmail.com")
    // check(err)

    // if err = client.Auth(auth); err != nil {
    //     fmt.Println(err)
    // }

    // if err = client.Mail(m.user); err != nil {
    //     fmt.Println(err)
    // }

    // if err = client.Rcpt(toId); err != nil {
    //     fmt.Println(err)
    // }

    // w, err := client.Data()
    // check(err)
    toIdarr:=[]string{toId}
    // msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", m.user, toId, title, text)
    msg := fmt.Sprintf("Subject: %s\r\n\r\n%s", title, text)
    err:=smtp.SendMail("smtp.gmail.com:587",auth,m.user,toIdarr,[]byte(msg))
    // _, err = w.Write([]byte(msg))
    if ok:=check(err);!ok{
        return false
    }
    // err = w.Close()
    return true
    // client.Quit()
}