package main

import (
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	dal.Init()

	//	CURD
	//mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "123456"})
	//mysql.DB.Model(&model.User{Email: "demo@example.com"}).Update("Password", "654321")

	//var row model.User
	//mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	//fmt.Println(row)

	//软删除
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Delete(&model.User{})

}
