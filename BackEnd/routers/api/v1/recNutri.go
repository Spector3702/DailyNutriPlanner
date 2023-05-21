package v1

import (
	. "BackEnd/models"
	service "BackEnd/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// RecNutriRouter a contract what this router can do
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
	recNutri.GET("/nutri/", r.GetNutriByGenderAndAge) // ex: localhost:8080/api/v1/recNutri/nutri?gender=male&age=19-44
	recNutri.POST("/analysis/", r.GetAnalysis)
	recNutri.GET("/calory/", r.GetCaloryByInfo)
	recNutri.GET("/food/", r.GetRandFoodBySpecfNutri)
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

// get analysis according to record and recNutri
func (r *recNutriRouter) GetAnalysis(c *gin.Context) {
	formFields := []string{"gender", "age", "calorie", "protein", "fat", "carbohydrate", "vitaminB1", "vitaminB2", "vitaminC", "nicotine", "vitaminB6", "vitaminA", "vitaminE", "ca", "p", "fe", "mg", "zn", "na", "k"}
	formValues := make([]string, len(formFields))

	for i, field := range formFields {
		formValues[i] = c.PostForm(field)
	}

	gender, age := formValues[0], formValues[1]
	nutritionStr := formValues[2:]

	analysis, err := service.RecNutriService.AnalysisByRecord(gender, age, nutritionStr...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"analysis": analysis})
}

// get calory based on all needed info
func (r *recNutriRouter) GetCaloryByInfo(c *gin.Context) {
	weightStr := c.Query("weight")
	heightStr := c.Query("height")
	workload := c.Query("workload")
	gender := c.Query("gender")
	age := c.Query("age")

	calory, err := service.RecNutriService.CalculateCaloryByInfo(weightStr, heightStr, workload, gender, age)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"calory": calory})
}

// get a random food that has enough specific nutri
func (r *recNutriRouter) GetRandFoodBySpecfNutri(c *gin.Context) {
	nutri := c.Query("nutri")
	need := c.Query("need")

	food, err := service.RecNutriService.GetFoodFromNutriTable(nutri, need)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"food": food})
}
