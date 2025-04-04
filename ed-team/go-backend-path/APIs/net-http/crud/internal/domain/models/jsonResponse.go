package models

type BaseJSONResponse struct {
	StatusCode int    `json:"status_code"`
	UUID       string `json:"uuid"`
	DateTime   string `json:"date_time"`
}

type ErrorResponse struct {
	BaseJSONResponse
	Error string `json:"error"`
}

type GenericResponse struct {
	BaseJSONResponse
	Description string `json:"description"`
	Data        any    `json:"data"`
}
