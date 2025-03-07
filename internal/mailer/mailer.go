package mailer

import "embed"

const (
	FromName              = "Go Ecommerce"
	UserWelcomeTemplate   = "user_invitation.tmpl"
	PasswordResetTemplate = "password_reset.tmpl"
)

//go:embed templates/*
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (string, error)
}
