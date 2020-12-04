package Routes

import (
	"extendvalidity-job/Controllers"
	"extendvalidity-job/Middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine  {
	r := gin.Default()
	r.Use(Middlewares.SignRequest)
	r.POST("/extend",Controllers.ExtendValidity)

	mockGroup := r.Group("/mock")
	{
		mockGroup.GET("upcc/:msisdn",Controllers.UpccMock)
		mockGroup.POST("subscriberinfo/",Controllers.SubscriberInfoMock)
		mockGroup.POST("validity/",Controllers.SetValidity)
	}

	return r
}
