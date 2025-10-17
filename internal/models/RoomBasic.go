package models

import (
	"time"

	"gorm.io/gorm"
)

type RoomBasic struct {
	gorm.Model
	Identity  string    `gorm:"column:identity;type:varchar(36);uniqueIndex;not null" json:"identity"`
	Name      string    `gorm:"column:name;type:varchar(36);not null" json:"name"`
	BeginTime time.Time `gorm:"column:begin_time;type:datetime" json:"begin_time"`
	EndTime   time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`
	CreateId  uint      `gorm:"column:create_id;type:int(11);not null" json:"create_id"` //id of the room creator
}

func (table *RoomBasic) TableName() string {
	return "room_basic"

}
