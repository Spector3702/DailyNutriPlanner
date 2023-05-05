package router

import (
	model "BackEnd/models"
	v1 "BackEnd/routers/api/v1"

	"github.com/gin-gonic/gin"
)

var (
	recNutriRoute v1.RecNutriRouter = v1.NewRecNutriRouter(model.NewRecNutriModel())
)

func Routes(app *gin.Engine) {
	apiRoute := app.Group("api")
	recNutriRoute.Setup(apiRoute)
}
