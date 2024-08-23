package models

import (
	"database/sql"
	"fmt"
	u "ldchat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string       `gorm:"type:varchar(256);" json:"name"`
	Password      string       `gorm:"type:varchar(256);" json:"password"`
	Phone         string       `gorm:"type:varchar(256);" json:"phone"`
	Email         string       `gorm:"type:varchar(256);" json:"email"`
	Identity      string       `gorm:"type:varchar(256);" json:"identity"`
	ClientIp      string       `gorm:"type:varchar(256);" json:"client_ip"`
	CLientPort    string       `gorm:"type:varchar(256);" json:"client_port"`
	LoginTime     sql.NullTime `json:"login_time"`
	LogoutTime    sql.NullTime `json:"logout_time"`
	HeartbeatTime sql.NullTime `json:"heartbeat_time"`
	IsLogout      bool         `json:"is_logout"`
	DeviceInfo    string       `gorm:"type:varchar(255);" json:"device_info"`
}

func GetUserList() ([]*UserBasic, error) {
	data := make([]*UserBasic, 10)
	u.MyDB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data, nil
}

func GetUserById(id int) (*UserBasic, error) {
	model := &UserBasic{}
	result := u.MyDB.Where("Id = ?", id).First(model)
	if result.Error == nil && result.RowsAffected > 0 {
		return model, nil
	} else {
		return nil, result.Error
	}
}
