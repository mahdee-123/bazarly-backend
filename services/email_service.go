package services

import (
	"fmt"
	"net/smtp"

	"github.com/mahdee-123/bazarly-backend/config"
)

func SendVerificationEmail(toEmail, name, token string) error {
	from := config.App.SMTPEmail
	password := config.App.SMTPPass
	host := config.App.SMTPHost
	port := config.App.SMTPPort

	verifyURL := fmt.Sprintf(
		"%s/api/sellers/verify-email?token=%s",
		config.App.AppURL,
		token,
	)

	subject := "Subject: Bazarly — Email Verification\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body := fmt.Sprintf(`
		<h2>Welcome to Bazarly, %s!</h2>
		<p>Please verify your email by clicking the link below:</p>
		<a href="%s" style="
			background:#534AB7;
			color:white;
			padding:12px 24px;
			border-radius:8px;
			text-decoration:none;
		">Verify Email</a>
		<p>This link expires in 24 hours.</p>
		<p>If you did not create an account, ignore this email.</p>
	`, name, verifyURL)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(
		host+":"+port,
		auth,
		from,
		[]string{toEmail},
		message,
	)

	return err
}