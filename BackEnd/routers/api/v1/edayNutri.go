package v1

import (
	. "BackEnd/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type EverydayNutriRouter struct {
	model EverydayNutriModel
}

func NewEverydayNutriRouter(m EverydayNutriModel) EverydayNutriRouter {
	return EverydayNutriRouter{
		model: m,
	}
}
func (r *EverydayNutriRouter) GetAllEveNutris(c *gin.Context) {
	eveNutris, err := r.model.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": eveNutris})
}
func (r *EverydayNutriRouter) Setup(rg *gin.RouterGroup) {
	everydayNutri := rg.Group("v1/everydayNutri")
	everydayNutri.GET("/", r.GetAllEveNutris)
}
