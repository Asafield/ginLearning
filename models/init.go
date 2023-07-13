package models

import (
	"fmt"
	"ginLearning/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	Database()
	DB.AutoMigrate(&SysUser{})
}
func Database() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//todo: logging
		fmt.Println(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		//todo logging
		fmt.Println("mysql lost: %v", err)
		panic(err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(100)
	DB = db
	fmt.Println("connect to database successfully!")

}
