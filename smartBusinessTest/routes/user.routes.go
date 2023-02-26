package routes

import (
	"github.com/gin-gonic/gin"

	"smartBusinessTest/controllers"
	// "smartBusinessTest/middleware"

)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	//router.Use(middleware.DeserializeUser())
	router.POST("/", uc.userController.AddUser)
	router.GET("/", uc.userController.GetUsers)
	router.PUT("/:userId", uc.userController.UpdateUser)
	router.GET("/:userId", uc.userController.GetUserById)
	router.DELETE("/:userId", uc.userController.DeleteUser)
}
