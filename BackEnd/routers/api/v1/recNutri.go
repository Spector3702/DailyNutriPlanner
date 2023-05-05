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
	recNutri.PUT("/:uid", r.UpdateUser)
	recNutri.DELETE("/:uid", r.DeleteUser)
	recNutri.GET("/:uid", r.GetUser)
}

// GetUsers 	 godoc
// @Summary      Get all users' information
// @Tags         recNutris
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.RecNutriModel
// @Failure      500  {object}  nil
// @Router       /recNutri [get]
func (r *recNutriRouter) GetAllRecNutris(c *gin.Context) {
	recNutri, _ := r.model.GetAll()
	c.IndentedJSON(http.StatusOK, recNutri)
}

// UpdateUser    godoc
// @Summary      update the user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param uid path int true "uid"
// @Param nickname formData string true "nickname"
// @Param password formData string true "password"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /users/:id [put]
func (r *recNutriRouter) UpdateUser(c *gin.Context) {
}

// DeleteUser    godoc
// @Summary      delete the user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param uid path int true "uid"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /users/:id [delete]
func (r *recNutriRouter) DeleteUser(c *gin.Context) {
}

// GetUser       godoc
// @Summary      Get the user's information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param uid path int false "uid"
// @Success      200  {object}  model.UserModel
// @Failure      500  {object}  nil
// @Router       /users/:id [get]
func (r *recNutriRouter) GetUser(c *gin.Context) {

}
