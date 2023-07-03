package config

import (
	"github.com/jinzhu/gorm"
	"github.com/rimo10/youtube-api-server/db"
)

var Database *gorm.DB

type Searchapi struct {
	gorm.Model
	Query       string `json:"Query"`
	Etag        string `json:"Etag"`
	VideoId     string `json:"VideoId"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	ChannelId   string `json:"ChannelId"`
	ChannelName string `json:"ChannelName"`
	PublishedAt string `json:"PublishedAt"`
}

func init() {
	db.Connect()
	Database = db.GetDB()
	if Database != nil {
		Database.AutoMigrate(&Searchapi{})
	}
}
