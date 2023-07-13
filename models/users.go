package models

import (
	"fmt"
	"time"
)

type SysUser struct {
	UserID       int       `gorm:"primary_key" json:"user_id"`
	UserName     string    `json:"username"`
	Password     string    `json:"password"`
	Salt         string    `json:"salt"`
	Email        string    `json:"email"`
	Mobile       string    `json:"mobile"`
	Company      string    `json:"company"`
	Status       string    `json:"status"`
	CreateUserID string    `json:"create_user_id"`
	GmtCreate    time.Time `json:"gmt_create"`
	GmtModified  time.Time `json:"gmt_modified"`
}

func UserSignin(username, encpwd string) bool {
	//var user SysUser
	fmt.Println("this is a test message")
	sysuser := SysUser{UserID: 1, UserName: "Asafield", Password: "admin"}
	err := DB.Create(&sysuser).Error
	if err != nil {
		fmt.Println("failed to create record: ", err)
	}

	return true
}
