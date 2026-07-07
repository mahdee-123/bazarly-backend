package services

import (
	"fmt"

	"github.com/mahdee-123/bazarly-backend/internal/config"
	"github.com/resend/resend-go/v3"
)

func SendVerificationEmail(toEmail, name, token string) error {
	client := resend.NewClient(config.App.ResendAPIKey)

	verifyURL := fmt.Sprintf(
		"%s/api/sellers/verify-email?token=%s",
		config.App.AppURL,
		token,
	)

	html := fmt.Sprintf(`
		<div style="font-family:sans-serif;max-width:480px;margin:0 auto;padding:32px">
			<h2 style="color:#1A1A1A;font-size:20px;font-weight:500">
				Welcome to Bazarly, %s!
			</h2>
			<p style="color:#6B6B6B;font-size:14px;line-height:1.6">
				Click the button below to verify your email.
			</p>
			<a href="%s" style="
				display:inline-block;
				margin-top:16px;
				background:#534AB7;
				color:white;
				padding:12px 24px;
				border-radius:8px;
				text-decoration:none;
				font-size:14px;
				font-weight:500;
			">Verify email</a>
			<p style="color:#ADADAA;font-size:12px;margin-top:24px">
				Link expires in 24 hours.
			</p>
		</div>
	`, name, verifyURL)

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To: []string{"arnobdev10@gmail.com"}, // সবসময় তোমার email এ যাবে testing এ,
		Subject: "Verify your Bazarly account",
		Html:    html,
	}

	_, err := client.Emails.Send(params)
	return err
}



func SendPasswordResetEmail(toEmail, name, token string) error {
	client := resend.NewClient(config.App.ResendAPIKey)

	resetURL := fmt.Sprintf(
		"%s/reset-password?token=%s",
		config.App.AppURL,
		token,
	)

	html := fmt.Sprintf(`
		<div style="font-family:sans-serif;max-width:480px;margin:0 auto;padding:32px">
			<h2 style="color:#1A1A1A;font-size:20px;font-weight:500">
				Reset your password
			</h2>
			<p style="color:#6B6B6B;font-size:14px;line-height:1.6">
				Hi %s, click the button below to reset your Bazarly password.
			</p>
			<a href="%s" style="
				display:inline-block;
				margin-top:16px;
				background:#534AB7;
				color:white;
				padding:12px 24px;
				border-radius:8px;
				text-decoration:none;
				font-size:14px;
				font-weight:500;
			">Reset password</a>
			<p style="color:#ADADAA;font-size:12px;margin-top:24px">
				This link expires in 1 hour. If you did not request a password reset, ignore this email.
			</p>
		</div>
	`, name, resetURL)

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{"arnobdev10@gmail.com"},
		Subject: "Reset your Bazarly password",
		Html:    html,
	}

	_, err := client.Emails.Send(params)
	return err
}