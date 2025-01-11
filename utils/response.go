package utils

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Responses struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	StatusCode int    `json:"message"`
	Message    string `json:"message"`
}
