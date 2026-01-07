package dto

type APIResponse struct {
	InternalCode int    `json:"internal_code"`
	Slug         string `json:"slug"`
	Message      string `json:"message"`
	Data         any    `json:"data,omitempty"`
}

type UploadPresignedURLResponse struct {
	Url string `json:"url"`
	Key string `json:"key"`
}
