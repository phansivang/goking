package utils

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Responses struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
}

type SuccessResponses struct {
	Message    string `json:"message"`
	StatusCode int    `json:"StatusCode"`
}

type ErrorResponse struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"message"`
}
