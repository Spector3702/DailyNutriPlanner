package v1

import (
	. "BackEnd/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// FoodstuffRouter a contract what this router can do
type FoodstuffRouter interface {
	Setup(rg *gin.RouterGroup)
}

type foodstuffRouter struct {
	model FoodstuffModel
}

func NewFoodstuffRouter(m FoodstuffModel) FoodstuffRouter {
	return &foodstuffRouter{
		model: m,
	}
}

func (r *foodstuffRouter) Setup(rg *gin.RouterGroup) {
	foodstuff := rg.Group("v1/foodstuff")
	//foodstuff.GET("/", r.GetAllFoodstuff)
	foodstuff.GET("/name/", r.GetFoodstuffByName)
	foodstuff.GET("/sno/", r.GetFoodstuffBySno)
}

// get foodstuff based on name from front-end
func (r *foodstuffRouter) GetFoodstuffByName(c *gin.Context) {
	name := c.Param("name")

	foodstuff, err := r.model.GetByName(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, foodstuff)
}

// get foodstuff based on Sno from front-end
func (r *foodstuffRouter) GetFoodstuffBySno(c *gin.Context) {
	sno := c.Param("sno")

	foodstuff, err := r.model.GetBySno(sno)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, foodstuff)
}
