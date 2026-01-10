package dto

type ForgotPasswordData struct {
	Email    string `json:"email"`
	Otp      string `json:"otp"`
	Attempts int    `json:"attempts"`
}

type AuthEmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Otp     string `json:"otp"`
}
