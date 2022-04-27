package main

import (
    "fmt"
  "net/smtp"
)

func sendemail(email string,) bool {
                // Sender data.
                from := "dev.codetuna@gmail.com"
                password := "qwertY3250"

                // Receiver email address.
                to := []string{
                 email,
                }

                // smtp server configuration.
                smtpHost := "smtp.gmail.com"
                smtpPort := "587"

                // Message.
                message := []byte("This is a test email message.")

                // Authentication.
                auth := smtp.PlainAuth("", from, password, smtpHost)

                // Sending email.
                err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
                if err != nil {
                fmt.Println(err)
                return false
                }
                fmt.Println("Email Sent Successfully!")
             return true
}
