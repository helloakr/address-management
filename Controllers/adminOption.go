//The file Controller file

package controllers

import (
	"hummingbird/databases"
	"hummingbird/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 管理员数据结构
type UserParameter struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
}

type ClassName struct {
	ClaName string `json:"className"`
}

//控制路由
func Login(c *gin.Context) {
	var usr UserParameter //用户参数结构体
	var user models.UserTable
	c.BindJSON(&usr) //解析请求转化为参数结构体

	DB := databases.GetDB()
	DB.Where(&models.UserTable{Name: usr.Name, Password: usr.Password}).Find(&user)
	if user.Name == usr.Name && user.Password == usr.Password {
		c.JSON(http.StatusOK, gin.H{"status": 1})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0})
	}

}

// // 控制路由
func Register(c *gin.Context) {
	var usr UserParameter
	c.BindJSON(&usr)
	DB := databases.GetDB()
	var user = models.UserTable{Name: usr.Name, Password: usr.Password}
	result := DB.Create(&user)

	if result != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0})
	}

}

func AddClass(c *gin.Context) {
	var clnNa ClassName
	c.BindJSON(&clnNa)
	time := time.Now().Format("2006-01-02")
	var cln = models.ClassName{Name: clnNa.ClaName, Time: time}
	DB := databases.GetDB()
	result := DB.Create(&cln)
	if result != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0})
	}
}
