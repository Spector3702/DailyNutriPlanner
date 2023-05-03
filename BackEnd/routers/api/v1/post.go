package v1

import "github.com/gin-gonic/gin"

type PostRouter interface {
	Setup(rg *gin.RouterGroup)
}

type postRouter struct {
}

func NewPostRouter() PostRouter {
	return &postRouter{}
}

func (r *postRouter) Setup(rg *gin.RouterGroup) {
	user := rg.Group("v1/posts")
	user.POST("/", r.CreatePost)
	user.GET("/", r.GetPosts)
	user.PUT("/:id", r.UpdatePost)
	user.DELETE("/:id", r.DeletePost)
	user.GET("/:id", r.GetPost)
}

// CreatePost 	 godoc
// @Summary      Create a post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param topic formData string true "topic"
// @Param authorId formData int true "authorId"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Success      201  {object}  nil
// @Failure      500  {object}  nil
// @Router       /posts [post]
func (r *postRouter) CreatePost(c *gin.Context) {

}

// GetPosts 	 godoc
// @Summary      Get all posts of the topic
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param topic formData string true "topic"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /posts [get]
func (r *postRouter) GetPosts(c *gin.Context) {

}

// UpdatePost 	 godoc
// @Summary      Update the post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Param title formData string false "title"
// @Param content formData string false "content"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /posts/:id [put]
func (r *postRouter) UpdatePost(c *gin.Context) {

}

// DeletePost 	 godoc
// @Summary      Delete the post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /posts/:id [delete]
func (r *postRouter) DeletePost(c *gin.Context) {

}

// GetPost 	     godoc
// @Summary      Get the post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /posts/:id [get]
func (r *postRouter) GetPost(c *gin.Context) {

}
