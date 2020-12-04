package Controllers

import (
	"extendvalidity-job/Services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
import "extendvalidity-job/Models/Dto"

func ExtendValidity(c *gin.Context)  {
	var req Dto.RequestExtend
	err := c.BindJSON(&req)
	if err!=nil{
		c.AbortWithStatus(http.StatusBadRequest)
	}else{
		log.Println("request: ",req)

		subinfo := Dto.SubscriberInfo{
			[]string{"EXP_DATE_MSISDN"},req.Msisdn,
		}

		rplPayload := Dto.RplRequest{
			"setEpiredDate",[]Dto.Param{
			{"msisdn",req.Msisdn},
			{"operation","setExpiredDate"},
			{"dateoffset","14"}},

		}

		if Services.GetDataSubscription(subinfo){
			c.AbortWithStatus(http.StatusBadRequest)
		}else if Services.GetUpccData(req.Msisdn){
			c.AbortWithStatus(http.StatusBadRequest)
		}else if Services.SetValidity(rplPayload){
			c.AbortWithStatus(http.StatusBadRequest)
		}else {
			c.JSON(http.StatusAccepted,
				Dto.ResponseEtend{http.StatusAccepted,"success",
					req})
		}


	}

}
