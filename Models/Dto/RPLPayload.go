package Dto

type RplRequest struct {
	Action string `json:"action"`
	Param []Param `json:"param"`
}

type Param struct {
	Paramname string `json:"paramname"`
	Paramvalue string `json:"paramvalue"`
}