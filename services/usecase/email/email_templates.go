package email

import (
	"fmt"
)

func SendEmailNotificationTemplate(name, email, content string) string {
	return fmt.Sprintf(`
    <html>
    <body style="font-family: Arial, sans-serif; color: #333;">
        <div style="max-width: 600px; margin: 0 auto; padding: 20px; background-color: #f4f4f4; border-radius: 8px;">
            <p>A user has contacted us with the following details:</p>
            <p>Name: <strong>%s</strong></p>
            <p>Email: <strong>%s</strong></p> 
            <p>Description:</p>
            <p style="background-color: #fff; padding: 15px; border-radius: 8px;">%s</p>
            <p>Best Regards,<br/>The TenderLinks Team</p>
            <hr style="border: 0; border-top: 1px solid #ddd; margin: 20px 0;">
            <p style="font-size: 12px; color: #888;">This email was sent from TenderLinks. If you did not send this message, please ignore this email.</p>
        </div>
    </body>
    </html>
    `, name, email, content)
}
