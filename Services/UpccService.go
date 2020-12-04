package Services

import (
	"bytes"
	"encoding/json"
	"extendvalidity-job/Models/Dto"
	"log"
	"net/http"
)

func GetUpccData(msisdn string)bool  {
	URL := "localhost:8080/mock/upcc/"+msisdn
	var rplResp Dto.UpccData
	response,err := http.Get(URL)
	if err != nil{
		log.Println("Error : ",err.Error())
		return false
	}else{
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		json.Unmarshal(buf.Bytes(),&rplResp)
		log.Println("upccdata : ",rplResp)
		return true
	}



}
