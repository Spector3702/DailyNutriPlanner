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
	eveNutris, err := r.model.GetAll() //analysis always show ""
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": eveNutris})
}

func (r *EverydayNutriRouter) Setup(rg *gin.RouterGroup) {
	everydayNutri := rg.Group("v1/everydayNutri")
	everydayNutri.GET("/", r.GetAllEveNutris)
	everydayNutri.POST("/GetNutri/", r.GetNutriByName)

}

func (r *EverydayNutriRouter) GetNutriByName(c *gin.Context) {
	name := c.Query("name")

	Nutri, err := r.model.GetByName(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, Nutri)
}

//
//"analysis_item":    Nutri.Analysis_item, //always show ""
//"Unit":             Nutri.Unit,
//"Content_per_unit": Nutri.Content_per_unit,
