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
	persnInfo.POST("/update/", r.UpdateByAllInfo)
}

// create new persnInfo tuple by all needed info
func (r *persnInfoRouter) RegisNewInfo(c *gin.Context) {
	email := c.PostForm("email")
	gender := c.PostForm("gender")
	age := c.PostForm("age")
	weightStr := c.PostForm("weight")
	heightStr := c.PostForm("height")
	work_load := c.PostForm("work_load")

	persnInfo, err := service.PersnInfoService.CreateByInfo(email, gender, age, weightStr, heightStr, work_load)
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
		"email":     persnInfo.Email,
		"gender":    persnInfo.Gender,
		"age":       persnInfo.Age,
		"weight":    persnInfo.Weight,
		"height":    persnInfo.Height,
		"work_load": persnInfo.Work_load,
	})
}

// update persnInfo by whole form from front end
func (r *persnInfoRouter) UpdateByAllInfo(c *gin.Context) {
	email := c.PostForm("email")
	gender := c.PostForm("gender")
	age := c.PostForm("age")
	weightStr := c.PostForm("weight")
	heightStr := c.PostForm("height")
	work_load := c.PostForm("work_load")

	persnInfo, err := service.PersnInfoService.UpdateByInfo(email, gender, age, weightStr, heightStr, work_load)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update fail"})
		return
	}

	c.JSON(201, persnInfo)
}
