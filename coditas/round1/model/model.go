package model

type PANDetails struct {
	Name   string `json:"name" validate:"required`
	Pan    string `json:"name" validate:"required,panValidator`
	Mobile int    `json:"name" validate:"required,mobileValidator`
	Email  string `json:"name" validate:"required,email`
}

type CustomResponse struct {
	StatusCode int
	Msg        interface{}
	Err        error
}
