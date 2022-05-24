package pojo

import (
	"GolangAPI/database"
)

//log的定義
type Log struct {
	Id     int    `json:"logId"` //定義json的name是什麼
	Ip     string `json:"logip"`
	Url    string `json:"logurl"`
	Status string `json:"logstatus"`
}

func FindALLlogs() []Log {
	var logs []Log
	database.DBconnect.Find(&logs)
	return logs
}

//因為c.Param("id")是string
func FindBylogId(logId string) Log {
	var log Log
	database.DBconnect.Where("id = ?", logId).First(&log)
	return log
}

//post   //要接收的值
func CreatLog(log Log) Log {
	database.DBconnect.Create(&log)
	return log //創建user並且回傳
}

/*
//delete user   //要接收的值是id
func DeleteUser(userId string) bool {
	user := User{}
	result := database.DBconnect.Where("id = ?", userId).Delete(&user)

	return result.RowsAffected != 0
	//刪除後會回傳此user
}

//update user
func UpdateUser(userId string, user User) User {

	database.DBconnect.Model(&user).Where("id = ?", userId).Updates(user) //Updates更新不只一處
	return user
}

*/
