package core

import "net/smtp"

func SendEmail(to string, subject string, body string) error {
	from := "erkidhoxholli@gmail.com"
	pass := "AchillesGod20"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}

	return nil
}
