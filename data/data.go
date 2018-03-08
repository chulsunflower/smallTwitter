package data

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name string
	Email string
	PhoneNumber string
	Password string
	Post []Post
}

type Post struct{
	gorm.Model
	UserID uint
	Content string
}



func init() {
	var err error
	DB, err = gorm.Open("mysql", "rli429:TKtk810626@tcp(reneedb.csu8todrdauu.ap-southeast-2.rds.amazonaws.com)/small_twitter?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error")
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Post{})
	return

}


func AddUser(name string, email string, phone string, password string) {
	passwordArray := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordArray, 10)
	if err !=nil{
		fmt.Print("Hash Error!")
	}
	pw := string(hash)

	DB.Create(&User{Name:name, Email:email, PhoneNumber:phone, Password:pw})

}


func CheckUser(email string, password string) bool{
	user := User{}
	DB.Where("email = ?",email).Find(&user)
	hashedPassword := []byte(user.Password)
	inputPassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, inputPassword)
	if err==nil{
		return true

	}
	return false

}

func SendPost(email string, content string)(err error){
	user := User{}
	DB.Where("email = ?", email).Find(&user)
	DB.Model(&user).Association("Post").Append(Post{Content:content})
	fmt.Printf("post content %s", user.Post)
	if err := DB.Model(&user).Association("Post").Append(Post{Content:content}).Error; err!=nil{
		return err
	}
	return
}

//有问题
func GetPost(email string)(err error){
	user := User{}
	post := Post{}
	DB.Where("email = ?", email).Find(&user)
	DB.Model(&user).Association("Post").Find(&post)
	if err := DB.Where("email = ?", email).Find(&user).Error; err!=nil{
		return err
	}
	return

}

func DeletePost(email string, content string)(err error){
	user := User{}
	fmt.Println(email, content)
	DB.Where("email = ?", email).Find(&user)
	DB.Where("user_id = ? and content = ?", user.ID, content).Delete(Post{})
	if err := DB.Where("user_id = ? and content = ?", user.ID, content).Delete(Post{}).Error; err!=nil{
		return err
	}
	return

}