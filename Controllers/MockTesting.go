package Controllers

import (
	"extendvalidity-job/Models/Dto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SubscriberInfoMock(c *gin.Context)  {
	var req Dto.SubscriberInfo
	c.BindJSON(&req)
	log.Println("SubsInfo req parsed : ",req)
	c.JSON(http.StatusOK,Dto.GenericResponse{
		http.StatusOK,"00","success",Dto.Param{
			"EXP_DATE","31-12-2020",
		},
	})
}

func UpccMock(c *gin.Context)  {

	msisdn:= c.Param("msisdn")
	log.Println("Upcc msisdn : ",msisdn)
	c.JSON(http.StatusOK,Dto.GenericResponse{
		http.StatusOK,"00","success",
		Dto.UpccData{

		},
	})
}

func SetValidity(c *gin.Context)  {

	msisdn:= c.Param("msisdn")
	log.Println("Upcc msisdn : ",msisdn)
	c.JSON(http.StatusOK,Dto.GenericResponse{
		http.StatusOK,"00","success",
		nil,
	})
}
