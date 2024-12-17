package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"gopkg.in/gomail.v2"
)

var opts struct {
	EmailAddress string `long:"email-addr"`
	EmailPort    int    `long:"email-port"`
	EmailCreds   string `long:"email-creds"`
	Insecure     bool   `long:"insecure"`
	Help         bool   `short:"h" long:"help"`
}

var help = `Test parameters:
--email-addr  - email client address
--email-port  - email client port
--email-creds - email "mail:password"
--insecure    - insecure connection`

func main() {
	_, err := flags.NewParser(&opts, flags.IgnoreUnknown).Parse()
	if err != nil {
		panic(err)
	}
	if opts.Help {
		fmt.Println(help)
		return
	}

	splt := strings.Split(opts.EmailCreds, ":")

	d := gomail.NewDialer(opts.EmailAddress, opts.EmailPort, splt[0], splt[1])
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: opts.Insecure,
		ServerName:         opts.EmailAddress,
	} //nolint:gosec

	mes := gomail.NewMessage()

	mes.SetHeader("From", splt[0])
	mes.SetHeader("To", splt[0])
	mes.SetHeader("Subject", "Email check!")
	mes.SetBody("text/html", "Hello from email checker.")

	err = d.DialAndSend(mes)
	if err != nil {
		fmt.Printf("unable to send email notification: %+v", err)
		os.Exit(1)
	}
	fmt.Println("checked successfully, email is ready.")
}
