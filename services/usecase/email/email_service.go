package email

import (
	"training-backend/package/config"
	"training-backend/package/log"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func NewEmailService(host string, port int, username, password string) *EmailService {
	return &EmailService{
		SMTPHost:     host,
		SMTPPort:     port,
		SMTPUsername: username,
		SMTPPassword: password,
	}
}

// SendEmail sends an email with the specified parameters
func (s *EmailService) SendEmail(to []string, subject, body string) error {
	m := gomail.NewMessage()

	// Set the sender
	m.SetHeader("From", s.SMTPUsername)

	// Set the recipients
	m.SetHeader("To", to...)

	// Set the email subject
	m.SetHeader("Subject", subject)

	// Set the body (HTML or plain text)
	m.SetBody("text/html", body)

	// Set up the SMTP dialer
	d := gomail.NewDialer(s.SMTPHost, s.SMTPPort, s.SMTPUsername, s.SMTPPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendEmailNotification(name, email, notificationType, content string) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	emailService := NewEmailService(cfg.Email.SMTPHost, cfg.Email.SMTPPort, cfg.Email.SMTPUser, cfg.Email.SMTPPassword)
	emailContent := SendEmailNotificationTemplate(name, email, content)

	err = emailService.SendEmail([]string{cfg.Email.SMTPUser}, notificationType, emailContent)
	if err != nil {
		log.Errorf("error sending email: ", err)
		return err
	}
	return nil
}
