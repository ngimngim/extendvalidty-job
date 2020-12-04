package Dto

type UpccData struct {
	Subscriber []map[string]string `json:"subscriber"`
	SubscribedService []map[string]string `json:"subscribedservice"`
	SubscriberQuota []map[string]string `json:"subscriberquota"`
	Info map[string]string `json:"info"`
	ResultCode int `json:"result_code"`
}