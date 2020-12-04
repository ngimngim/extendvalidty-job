package Dto

type GenericResponse struct {
	Status int `json:"status"`
	Code string `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
