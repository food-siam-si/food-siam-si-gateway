package dto

type HelloWorldQuery struct {
	Text string `query:"text" validate:"required,min=3,max=10"`
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}
