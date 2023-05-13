package router

import (
	model "BackEnd/models"
	v1 "BackEnd/routers/api/v1"

	"github.com/gin-gonic/gin"
)

var (
	recNutriRoute  v1.RecNutriRouter      = v1.NewRecNutriRouter(model.NewRecNutriModel())
	persnInfoRoute v1.PersnInfoRouter     = v1.NewPersnInfoRouter(model.NewPersnInfoModel())
	foodstuffRoute v1.FoodstuffRouter     = v1.NewFoodstuffRouter(model.NewFoodstuffModel())
	eveNutrisRoute v1.EverydayNutriRouter = v1.NewEverydayNutriRouter(model.NewEverydayNutriModel())
)

func Routes(app *gin.Engine) {
	apiRoute := app.Group("api")
	recNutriRoute.Setup(apiRoute)
	persnInfoRoute.Setup(apiRoute)
	foodstuffRoute.Setup(apiRoute)
	eveNutrisRoute.Setup(apiRoute)
}
