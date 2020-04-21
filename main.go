package main

import (
	"fmt"
	"golang-blog/config"
	"golang-blog/db"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func SendMail(username, password, host, to, name, subject, body, mailType string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", username, password, hp[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + name + "<" + username + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, username, sendTo, msg)
	return err
}

func main() {

	time.Sleep(time.Second * 30)

	var (
		err error
	)

	if err = config.InitConfig(); err != nil {
		goto ERR
	}
	if err = db.InitRedisClient(); err != nil {
		goto ERR
	}

	//if _, err = db.G_redisClient.ZAdd("send:10",  &redis.Z{Score: float64(1), Member: 2}).Result(); err != nil {
	//	goto ERR
	//}
	InitFeedMgr()
	fmt.Println("success")
ERR:
	fmt.Fprint(os.Stderr, err)
	os.Exit(-1)
}
