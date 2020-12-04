package Controllers

import (
	"extendvalidity-job/Models/Dto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SubscriberInfoMock(c gin.Context)  {
	var req Dto.SubscriberInfo
	c.BindJSON(&req)
	log.Println("SubsInfo req parsed : ",req)
	c.JSON(http.StatusOK,Dto.GenericResponse{
		http.StatusOK,"00","success"
	})
}
