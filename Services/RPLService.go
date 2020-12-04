package Services

import (
	"bytes"
	"encoding/json"
	"extendvalidity-job/Models/Dto"
	"log"
	"net/http"
)

func SetValidity(payload Dto.RplRequest) bool  {
	URL := "http://extendvalidity-job.herokuapp.com/mock/validity/"
	var rplResp Dto.GenericResponse
	body,_ := json.Marshal(payload)
	ioBody:=bytes.NewReader(body)
	response,err := http.Post(URL,"application/json",ioBody)
	if err != nil{
		log.Println("Error : ",err.Error())
		return false
	}else{
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		json.Unmarshal(buf.Bytes(),&rplResp)
		log.Println("subsinfo : ",rplResp)
		return true
	}


}
