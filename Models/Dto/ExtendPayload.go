package Dto

type RequestExtend struct {
	ID int64 `json:"id"`
	Msisdn string `json:"msisdn"`
	SubscriberNo string `json:"subscriberno"`
}

type ResponseEtend struct {
	Code int `json:"code`
	Message string `json:"message"`
	Data RequestExtend `json:"data"`
	NewExpDate string `json:"new_exp_date"`
}
