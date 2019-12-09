package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id   int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Age  int    `gorm:"size:11;DEFAULT NULL" json:"age"`
	Name string `gorm:"size:255;DEFAULT NULL" json:"name"`
	//gorm后添加约束，json后为对应mysql里的字段
}

func main() {
	MysqlDB, err := gorm.Open("mysql", "root:123456@(192.144.99.100:3301)/test?charset=utf8")
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	} else {
		fmt.Println("connect database success")
		MysqlDB.SingularTable(true)
		MysqlDB.AutoMigrate(&User{}) //自动建表
		fmt.Println("create table success")
	}

	user := User{Id: 1, Age: 30, Name: "bingo"}
	MysqlDB.Create(&user) //创建对象

	defer MysqlDB.Close()
}
