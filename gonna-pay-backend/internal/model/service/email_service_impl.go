package service

import (
	"context"
	"fmt"
	"os"

	"github.com/resend/resend-go/v4"
)

type EmailService interface {
	SendVerificationEmail(ctx context.Context, toEmail, userName, token string) error
	SendPasswordResetEmail(ctx context.Context, toEmail, userName, token string) error
}

type emailService struct {
	client *resend.Client
	from   string
	appURL string
}

func NewEmailService() EmailService {
	return &emailService{
		client: resend.NewClient(os.Getenv("RESEND_API_KEY")),
		from:   os.Getenv("RESEND_FROM"),
		appURL: os.Getenv("APP_URL"),
	}
}

func (s *emailService) SendVerificationEmail(ctx context.Context, toEmail, userName, token string) error {
	link := fmt.Sprintf("%s/verify-email?token=%s", s.appURL, token)
	_, err := s.client.Emails.Send(&resend.SendEmailRequest{
		From:    s.from,
		To:      []string{toEmail},
		Subject: "Verify your email",
		Html:    fmt.Sprintf("<p>Hi %s, click <a href=\"%s\">here</a> to verify your email.</p>", userName, link),
	})
	return err
}

func (s *emailService) SendPasswordResetEmail(ctx context.Context, toEmail, userName, token string) error {
	link := fmt.Sprintf("%s/reset-password?token=%s", s.appURL, token)
	_, err := s.client.Emails.Send(&resend.SendEmailRequest{
		From:    s.from,
		To:      []string{toEmail},
		Subject: "Reset your password",
		Html:    fmt.Sprintf("<p>Hi %s, click <a href=\"%s\">here</a> to reset your password.</p>", userName, link),
	})
	return err
}
