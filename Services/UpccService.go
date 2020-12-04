package Services

import (
	"bytes"
	"encoding/json"
	"extendvalidity-job/Models/Dto"
	"log"
	"net/http"
)

func GetUpccData(msisdn string, res chan bool)  {
	defer close(res)
	URL := "http://extendvalidity-job.herokuapp.com/mock/upcc/"+msisdn
	var rplResp Dto.UpccData
	response,err := http.Get(URL)
	if err != nil{
		log.Println("Error : ",err.Error())
		res<- false
	}else{
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		json.Unmarshal(buf.Bytes(),&rplResp)
		log.Println("upccdata : ",rplResp)
		res<- true
	}



}
