package v1

import (
	. "BackEnd/models"
	service "BackEnd/services"
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
	user := rg.Group("v1/users")
	user.POST("/", r.CreateUser)
	user.GET("/", r.GetUsers)
	user.PUT("/:uid", r.UpdateUser)
	user.DELETE("/:uid", r.DeleteUser)
	user.GET("/:uid", r.GetUser)
}

// CreateUser 	 godoc
// @Summary      Register a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param nickname formData string true "nickname"
// @Param email formData string true "email"
// @Param password formData string true "password"
// @Success      201  {object}  nil
// @Failure      500  {object}  nil
// @Router       /users [post]
func (r *recNutriRouter) CreateUser(c *gin.Context) {
	gender := c.PostForm("gender")
	age := c.PostForm("age")
	// headshot := c.DefaultPostForm("headshot", "")
	user, err := service.RecNutriService.CreateUser(gender, age)
	if err != nil {
		c.JSON(http.StatusConflict,
			gin.H{"message": "Failed to create user"})
	}

	c.JSON(201, user)
}

// GetUsers 	 godoc
// @Summary      Get all users' information
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.UserModel
// @Failure      500  {object}  nil
// @Router       /users [get]
func (r *recNutriRouter) GetUsers(c *gin.Context) {
	users, _ := r.model.GetAll()
	c.IndentedJSON(http.StatusOK, users)
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
