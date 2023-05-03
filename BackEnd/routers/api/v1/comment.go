package v1

import "github.com/gin-gonic/gin"

type CommentRouter interface {
	Setup(rg *gin.RouterGroup)
}

type commentRouter struct {
}

func NewCommentRouter() CommentRouter {
	return &commentRouter{}
}

func (r *commentRouter) Setup(rg *gin.RouterGroup) {
	user := rg.Group("v1/comments")
	user.POST("/", r.CreateComment)
	user.GET("/", r.GetComments)
	user.PUT("/:id", r.UpdateComment)
	user.DELETE("/:id", r.DeleteComment)
	user.GET("/:id", r.GetComment)
}

// CreateComment godoc
// @Summary      Create a comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param postId formData int true "postId"
// @Param authorId formData int true "authorId"
// @Param content formData string true "content"
// @Success      201  {object}  nil
// @Failure      500  {object}  nil
// @Router       /comments [post]
func (r *commentRouter) CreateComment(c *gin.Context) {

}

// GetComments   godoc
// @Summary      Get all comments of the post
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param postId formData int true "postId"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /comments [get]
func (r *commentRouter) GetComments(c *gin.Context) {

}

// UpdateComment godoc
// @Summary      Update the comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Param content formData string true "content"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /comments/:id [put]
func (r *commentRouter) UpdateComment(c *gin.Context) {

}

// DeleteComment godoc
// @Summary      Delete the comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /comments/:id [delete]
func (r *commentRouter) DeleteComment(c *gin.Context) {

}

// GetComment    godoc
// @Summary      Get a comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      201  {object}  nil
// @Failure      500  {object}  nil
// @Router       /comments/:id [get]
func (r *commentRouter) GetComment(c *gin.Context) {

}
