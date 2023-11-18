package models

type CompileBody struct {
	ID string `json:"client_id"`
	Language string `json:"language"`
	Code string `json:"code"`
}

