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
			{"",""},req.Msisdn,
		}

		if Services.GetDataSubscription(subinfo){

		}else if Services.GetUpccData(req.Msisdn){

		}else if

		c.JSON(http.StatusAccepted,
			Dto.ResponseEtend{http.StatusAccepted,"success",
			req})
	}

}
