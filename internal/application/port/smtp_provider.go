package port

type SMTPProvider interface {
	Send(to, subject, body string) error
	
	AuthEmail(to, subject, otp string) error
}