package dto

type DTOError struct {
	Message string `json:"message"`
}

type DTOErrorArray struct {
	Message []string `json:"message"`
}

type DTOErrorWithCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
