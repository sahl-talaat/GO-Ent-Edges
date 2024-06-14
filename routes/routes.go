package routes

import (
	"test/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("users", controller.GetUsers)
		v1.POST("users", controller.CreateUser)
		v1.GET("users/:id", controller.GetUser)
		v1.PUT("users/:id", controller.UpdateUser)
		v1.DELETE("users/:id", controller.DeleteUser)
		v1.POST("/user/:userID/spouse/:spouseID", controller.SetSpouse)

		v1.GET("cards", controller.GetCards)
		v1.POST("cards/:ownerid", controller.CreateCard)
		v1.GET("cards/:id", controller.GetCard) /*
			v1.PUT("cards/:id", controller.Upda)
			v1.DELETE("cards/:id", controller.DeleteUser) */

		v1.GET("nodes", controller.GetNodes)
		v1.POST("nodes", controller.CreateNode)
		v1.GET("nodes/:id", controller.GetNode)

		v1.GET("pets", controller.GetPets)
		v1.POST("pets/:ownerID", controller.CreatePet)
		v1.GET("pets/:id", controller.GetPet)

		v1.POST("tree", controller.CreateTreeItem)
		v1.GET("tree/:id", controller.GetTreeItem)
		v1.POST("tree/:nodeID/:parentID", controller.SetParent)
		v1.GET("tree/leaf", controller.GetLeafNodes)
	}
}
