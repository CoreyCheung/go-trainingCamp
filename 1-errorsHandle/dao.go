package main

import (
	"github.com/jinzhu/gorm"
)

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
var (
	connString = "root@tcp(127.0.0.1:3306)/test"
	gormClient *gorm.DB
)

type User struct {
	Name string
	Age  int
}

//从业务逻辑上讲  找不到记录并不是一条错误,但是上层应该知道
func GetUserList(id string) (User, error) {
	var user User
	err := gormClient.Where("where id = ?", id).Find(user).Error
	return user,err

}
func init() {
	//这里省略超时重连等处理
	db, _ := gorm.Open("mysql", connString)
	gormClient = db
}
