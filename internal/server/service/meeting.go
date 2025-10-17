package service

import (
	"hermes/internal/helper"
	"hermes/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MeetingCreate(c *gin.Context) {
	in := new(MeetingCreateRequest)

	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "parameter err",
		})
	}
	err = models.DB.Create(&models.RoomBasic{
		Identity:  helper.GetUUID(),
		Name:      in.Name,
		BeginTime: time.UnixMilli(in.CreateAt),
		EndTime:   time.UnixMilli(in.DeleteAt),
		CreateId:  0, //todo:从auth中间件中获取用户的信息
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "DB create err" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "create success",
	})
}
