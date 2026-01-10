package constants

const (
	CodeSuccess                     = 1000
	CodeLoginSuccess                = 1001
	CodeLogoutSuccess               = 1002
	CodeChangePasswordSuccess       = 1003
	CodeForgotPasswordSuccess       = 1004
	CodeVerifyForgotPasswordSuccess = 1005
	CodeResetPasswordSuccess        = 1006
	CodeBadRequest                  = 4000
	CodeLoginFailed                 = 4001
	CodeInvalidToken                = 4002
	CodeUnAuth                      = 4003
	CodeNoRefreshToken              = 4004
	CodeUserNotFound                = 4005
	CodeInvalidPassword             = 4006
	CodeEmailDoesNotExist           = 4007
	CodeTooManyAttempts             = 4008
	CodeInvalidOTP                  = 4009
	CodeInternalError               = 5000

	ExchangeEmail       = "email.send"
	QueueNameAuthEmail  = "email.send.auth"
	RoutingKeyAuthEmail = "email.send.auth"
)
