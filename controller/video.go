package controller

import (
	"ByteGopher_SimpleDouyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type VideoController interface {
	Feed(c *gin.Context)

}

type videoController struct {
		db *gorm.DB
		service service.VideoService 
}


func NewVideoController(db *gorm.DB) VideoController {
	return &videoController{db: db, service: service.NewVideoService(db)}
}

// Feed same demo video list for every request
func (controller *videoController)Feed(c *gin.Context) {
	// latest_time, ok := c.GetQuery("latest_time")
	// token, ok := c.GetQuery("token")

	// book, err := controller.service.FindByID(c.Param("id"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }
	// return c.JSON(http.StatusOK, book)
	feed, err := controller.service.GetFeed()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, feed)
}
