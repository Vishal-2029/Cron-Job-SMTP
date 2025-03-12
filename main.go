package main

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	senderEmail := "vishal.9dot@gmail.com"
	senderPassword := "hnll ucfe nitp awux"

	receiverEmail := "vishsabhadiya002@gmail.com"

	subject := "Cron job email"
	body := "Using cron job to send email"
	message := []byte(subject + "\n" + "\n" + body)

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{receiverEmail}, message)
		if err != nil {
			fmt.Println("Error sending email:", err)
			return
		}
		fmt.Println("Email sent successfully!")
	})
	c.Start()

	time.Sleep(1 * time.Minute)
	fmt.Println("Stopping cron jobs...")
	defer c.Stop()
}
