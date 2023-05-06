package v1

import (
	. "BackEnd/models"
	service "BackEnd/services"

	// service "BackEnd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersnInfoRouter interface {
	Setup(rg *gin.RouterGroup)
}

type persnInfoRouter struct {
	model PersnInfoModel
}

func NewPersnInfoRouter(m PersnInfoModel) PersnInfoRouter {
	return &persnInfoRouter{
		model: m,
	}
}

func (r *persnInfoRouter) Setup(rg *gin.RouterGroup) {
	persnInfo := rg.Group("v1/persnInfo")
	persnInfo.POST("/", r.RegisNewInfo)
	persnInfo.GET("/", r.LoginByEmail)
}

// create new persnInfo tuple by all needed info
func (r *persnInfoRouter) RegisNewInfo(c *gin.Context) {
	email := c.PostForm("email")
	gender := c.PostForm("gender")
	age := c.PostForm("age")
	weightStr := c.PostForm("weight")
	heightStr := c.PostForm("height")
	workload := c.PostForm("workload")

	persnInfo, err := service.PersnInfoService.CreateByInfo(email, gender, age, weightStr, heightStr, workload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create fail"})
		return
	}

	c.JSON(201, persnInfo)
}

// get persnInfo tuple by given email
func (r *persnInfoRouter) LoginByEmail(c *gin.Context) {
	email := c.Query("email")

	persnInfo, err := r.model.GetByEmail(email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"Email":    persnInfo.Email,
		"Gender":   persnInfo.Gender,
		"Age":      persnInfo.Age,
		"Weight":   persnInfo.Weight,
		"Height":   persnInfo.Height,
		"WorkLoad": persnInfo.WorkLoad,
	})
}
