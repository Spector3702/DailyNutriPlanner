package v1

import (
	. "BackEnd/models"

	// service "BackEnd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRouter a contract what this router can do
type RecNutriRouter interface {
	Setup(rg *gin.RouterGroup)
}

type recNutriRouter struct {
	model RecNutriModel
}

func NewRecNutriRouter(m RecNutriModel) RecNutriRouter {
	return &recNutriRouter{
		model: m,
	}
}

func (r *recNutriRouter) Setup(rg *gin.RouterGroup) {
	recNutri := rg.Group("v1/recNutri")
	recNutri.GET("/", r.GetAllRecNutris)
	recNutri.GET("/nutri", r.GetNutriByGenderAndAge) // ex: localhost:8080/api/v1/recNutri/nutri?gender=male&age=19-44
}

// get all recNutri data
func (r *recNutriRouter) GetAllRecNutris(c *gin.Context) {
	recNutri, _ := r.model.GetAll()
	c.IndentedJSON(http.StatusOK, recNutri)
}

// get nutritions based on gender & age from front-end
func (r *recNutriRouter) GetNutriByGenderAndAge(c *gin.Context) {
	gender := c.Query("gender")
	age := c.Query("age")

	recommendedNutrition, err := r.model.GetByGenderAndAge(gender, age)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"Gender":          gender,
		"Age":             age,
		"Calorie":         recommendedNutrition.Calorie,
		"Protein":         recommendedNutrition.Protein,
		"Fat":             recommendedNutrition.Fat,
		"Carbohydrate":    recommendedNutrition.Carbohydrate,
		"VitaminB1":       recommendedNutrition.VitaminB1,
		"VitaminB2":       recommendedNutrition.VitaminB2,
		"VitaminC":        recommendedNutrition.VitaminC,
		"Nicotine":        recommendedNutrition.Nicotine,
		"VitaminB6":       recommendedNutrition.VitaminB6,
		"VitaminA":        recommendedNutrition.VitaminA,
		"VitaminE":        recommendedNutrition.VitaminE,
		"Ca":              recommendedNutrition.Ca,
		"P":               recommendedNutrition.P,
		"Fe":              recommendedNutrition.Fe,
		"Mg":              recommendedNutrition.Mg,
		"Zn":              recommendedNutrition.Zn,
		"Na":              recommendedNutrition.Na,
		"K":               recommendedNutrition.K,
		"NumbersOfPeople": recommendedNutrition.NumbersOfPeople,
	})
}
