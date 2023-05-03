package v1

import "github.com/gin-gonic/gin"

type TopicRouter interface {
	Setup(rg *gin.RouterGroup)
}

type topicRouter struct {
}

func NewTopicRouter() TopicRouter {
	return &topicRouter{}
}

func (r *topicRouter) Setup(rg *gin.RouterGroup) {
	topic := rg.Group("v1/topics")
	topic.POST("/", r.CreateTopic)
	topic.GET("/", r.GetTopics)
	topic.PUT("/:id", r.UpdateTopic)
	topic.DELETE("/:id", r.DeleteTopic)
	topic.GET("/:id", r.GetTopic)
}

// CreateTopic 	 godoc
// @Summary      Create a topic
// @Tags         topics
// @Accept       json
// @Produce      json
// @Param name formData string true "name"
// @Param description formData string true "description"
// @Success      201  {object}  nil
// @Failure      500  {object}  nil
// @Router       /topics [post]
func (r *topicRouter) CreateTopic(c *gin.Context) {
}

// GetTopics 	 godoc
// @Summary      Get all topics
// @Tags         topics
// @Accept       json
// @Produce      json
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /topics [get]
func (r *topicRouter) GetTopics(c *gin.Context) {

}

// UpdateTopic 	 godoc
// @Summary      Update the topic
// @Tags         topics
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Param name formData string false "name"
// @Param description formData string false "description"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /topics/:id [put]
func (r *topicRouter) UpdateTopic(c *gin.Context) {

}

// DeleteTopic 	 godoc
// @Summary      Delete the topic
// @Tags         topics
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /topics/:id [delete]
func (r *topicRouter) DeleteTopic(c *gin.Context) {

}

// GetTopic 	 godoc
// @Summary      Get the topic
// @Tags         topics
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      200  {object}  nil
// @Failure      500  {object}  nil
// @Router       /topics/:id [get]
func (r *topicRouter) GetTopic(c *gin.Context) {

}
