package pojo

import (
	"GolangAPI/database"
)

//user的定義
type User struct {
	Id       int    `json:"UserId"` //定義json的name是什麼
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func FindALLUsers() []User {
	var users []User
	database.DBconnect.Find(&users)
	return users
}

//因為c.Param("id")是string
func FindByUserId(userId string) User {
	var user User
	database.DBconnect.Where("id = ?", userId).First(&user)
	return user
}

//post   //要接收的值
func CreatUser(user User) User {
	database.DBconnect.Create(&user)
	return user //創建user並且回傳
}

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
