package mailer

import (
	"bytes"
	"errors"
	"html/template"

	"github.com/resend/resend-go/v2"
)

type ResendClient struct {
	fromEmail string
	apiKey    string
}

func NewResendClient(fromEmail, apiKey string) (ResendClient, error) {
	if apiKey == "" {
		return ResendClient{}, errors.New("resend api key is required")
	}
	return ResendClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}

func (r ResendClient) Send(templateFile, username, email string, data any, isSandbox bool) (string, error) {
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return "", err
	}
	// Render the subject
	subjectBuf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subjectBuf, "subject", data); err != nil {
		return "", err
	}
	// Render the body
	bodyBuf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(bodyBuf, "body", data); err != nil {
		return "", err
	}
	// Initialize the Resend Client
	client := resend.NewClient(r.apiKey)

	params := &resend.SendEmailRequest{
		From:    r.fromEmail,
		To:      []string{email},
		Subject: subjectBuf.String(),
		Html:    bodyBuf.String(),
	}
	sent, err := client.Emails.Send(params)
	if err != nil {
		return "", err
	}

	return sent.Id, nil
}
