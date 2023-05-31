package v1

import (
	"BackEnd/entity"
	model "BackEnd/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RecordRouter struct {
	r model.RecordModel
	f model.FoodstuffModel
}

var RecNutriService = NewRecordRouter()

func NewRecordRouter() *RecordRouter {
	return &RecordRouter{
		r: model.NewRecordModel(),
		f: model.NewFoodstuffModel(),
	}
}

func (r *RecordRouter) CreateRecord(c *gin.Context) {
	email := c.PostForm("email")
	name := c.PostForm("name")
	sno, err := r.f.GetSnoByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	currentTime := time.Now()
	record := &entity.Record{
		Email: email,
		Name:  name,
		Sno:   sno,
		Date:  currentTime,
	}

	err = r.r.CreateRecord(record)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record created successfully"})
}

func (r *RecordRouter) Setup(rg *gin.RouterGroup) {
	record := rg.Group("v1/record")
	record.POST("/create/", r.CreateRecord)
}
