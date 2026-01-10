package dto

type UploadPresignedURLRequest struct {
	FileName    string `json:"file_name" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
}

type UploadPresignedURLsRequest struct {
	Files []UploadPresignedURLRequest `json:"files" binding:"required,min=1,dive"`
}

type ViewPresignedURLsRequest struct {
	Keys []string `json:"keys" binding:"required,min=1,dive"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=6"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyForgotPasswordRequest struct {
	ForgotPasswordToken string `json:"forgot_password_token" binding:"required,uuid4"`
	Otp                 string `json:"otp" binding:"required,len=6,numeric"`
}

type ResetPasswordRequest struct {
	ResetPasswordToken string `json:"reset_password_token" binding:"required,uuid4"`
	NewPassword        string `json:"new_password" binding:"required,min=6"`
}
