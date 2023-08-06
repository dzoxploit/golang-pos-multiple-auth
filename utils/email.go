package utils

import (
	"net/smtp"
)

// SendEmail sends an email to the specified recipient with the given subject and body.
func SendEmail(recipientEmail, subject, body string) error {
	from := "didinnuryahya@gmail.com" // Replace with your email address
	password := "your-password" // Replace with your email password
	to := []string{recipientEmail}
	msg := []byte("To: " + recipientEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, to, msg)

	return err
}
