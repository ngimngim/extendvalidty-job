package Services

import (
	"bytes"
	"encoding/json"
	"extendvalidity-job/Models/Dto"
	"log"
	"net/http"
)

func GetDataSubscription(payload Dto.SubscriberInfo)bool  {
	URL := "localhost:8080/mock/subscriberinfo/"
	var subsinfo Dto.GenericResponse
	body,_ := json.Marshal(payload)
	ioBody:=bytes.NewReader(body)
	response,err := http.Post(URL,"application/json",ioBody)
	if err != nil{
		log.Println("Error : ",err.Error())
		return false
	}else{
		buf := new(bytes.Buffer)
	    buf.ReadFrom(response.Body)
		json.Unmarshal(buf.Bytes(),&subsinfo)
		log.Println("subsinfo : ",subsinfo)
		return  true
	}


}
