package app

import (
	"privy-test/usecases"

	"privy-test/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app usecases.IUsecasesInterface) {
	cakeSrv := handlers.NewHttpHandler(app)
	cake := r.Group("/cakes")
	{
		cake.POST("/", cakeSrv.PostCake)
		cake.GET("/", cakeSrv.GetAllCakes)
		cake.GET("/:id", cakeSrv.GetCakeDetailByID)
		cake.PATCH("/:id", cakeSrv.UpdateCakeByID)
		cake.DELETE("/:id", cakeSrv.DeleteCakeByID)
	}
}
