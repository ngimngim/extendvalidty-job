package Services

import (
	"bytes"
	"encoding/json"
	"extendvalidity-job/Models/Dto"
	"log"
	"net/http"
)

func GetDataSubscription(payload Dto.SubscriberInfo, res chan bool)  {
	defer close(res)
	URL := "http://extendvalidity-job.herokuapp.com/mock/subscriberinfo/"
	log.Println("subsinfo request payload : ",payload)
	var subsinfo Dto.GenericResponse
	body,_ := json.Marshal(payload)
	ioBody:=bytes.NewReader(body)
	response,err := http.Post(URL,"application/json",ioBody)
	if err != nil{
		log.Println("Error : ",err.Error())
		res<- false
	}else{
		buf := new(bytes.Buffer)
	    buf.ReadFrom(response.Body)
		json.Unmarshal(buf.Bytes(),&subsinfo)
		log.Println("subsinfo : ",subsinfo)

		if subsinfo.Code=="00"{
			res<-  true
		}else{
			res<- true
		}


	}


}
