package controller

import (
	"ByteGopher_SimpleDouyin/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Feed(c *gin.Context)

}

type videoController struct {
		service service.VideoService 
}


func NewVideoController() VideoController {
	return &videoController{service: service.NewVideoService()}
}

// Feed same demo video list for every request
func (controller *videoController)Feed(c *gin.Context) {
	// latest_time, ok := c.GetQuery("latest_time")
	// token, ok := c.GetQuery("token")
	feed, err := controller.service.GetFeed()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(feed.VideoList)
	c.JSON(http.StatusOK, feed)
}
