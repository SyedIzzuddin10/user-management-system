package dto

type Display struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorDisplay struct {
}
