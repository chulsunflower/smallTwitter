package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"smallTwitter/data"
)

//POST
func signup(writer http.ResponseWriter, request *http.Request){
	len := request.ContentLength
	body := make([]byte, len)
	request.Body.Read(body)
	var user data.User
	json.Unmarshal(body, &user)
	data.AddUser(user.Name, user.Email, user.PhoneNumber, user.Password)
	writer.WriteHeader(200)
}

//POST
func login(writer http.ResponseWriter, request *http.Request){
	len := request.ContentLength
	body := make([]byte, len)
	request.Body.Read(body)
	var user data.User
	json.Unmarshal(body, &user)
	loginSuccess := data.CheckUser(user.Email, user.Password)
	if loginSuccess{
		//cookie := http.Cookie{Name:"User_Cookie", Value:user.Email,HttpOnly:true}
		//http.SetCookie(writer, &cookie)
		//fmt.Println(cookie)
		//fmt.Println(writer.Header())
		fmt.Println("Login succesful!")
	}else{
		fmt.Println("Login Failed. Please try again!")
		writer.WriteHeader(400)
	}


}


