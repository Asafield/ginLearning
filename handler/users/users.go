package users

import (
	"encoding/json"
	"fmt"
	"ginLearning/models"
	"ginLearning/util"
	"github.com/gin-gonic/gin"
)

const (
	pwd_salt = "YzcmCZNvbXocrsz9dm8e"
)

func SignInHandler(c *gin.Context) {
	fmt.Println("-------dfdfs------------")
	var user map[string]string
	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &user)
	username := user["username"]
	password := user["password"]
	enc_passwd := util.MD5([]byte(password + pwd_salt))
	pwdChecked := models.UserSignin(username, enc_passwd)
	if !pwdChecked {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "用户名或密码错误",
		})
		return
	}
	msg := make(map[string]interface{})
	token, err := util.GenerateToken(username, enc_passwd)
	if err != nil {
		//todo: logging
		fmt.Println("failed to create token!")
		return
	}
	msg["expire"] = 43200
	msg["token"] = token
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  msg,
	})
}
