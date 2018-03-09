package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"smallTwitter/data"
)

//POST http://localhost:8080/signup
func signup(writer http.ResponseWriter, request *http.Request){
	len := request.ContentLength
	body := make([]byte, len)
	request.Body.Read(body)
	var user data.User
	json.Unmarshal(body, &user)
	err := data.AddUser(user.Name, user.Email, user.PhoneNumber, user.Password)
	if err != nil{
		errMessage := "Invalid email address, try again!"
		jsData, err := json.Marshal(errMessage)
		if err != nil{
			fmt.Println("Json Marshal Error!")
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsData)
		writer.WriteHeader(400)
	}else{
		successMessage := "Successfully registered!"
		jsData, err := json.Marshal(successMessage)
		if err != nil{
			fmt.Println("Json Marshal Error!")
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsData)
		writer.WriteHeader(200)
	}

}

//POST http://localhost:8080/login
func login(writer http.ResponseWriter, request *http.Request){
	len := request.ContentLength
	body := make([]byte, len)
	request.Body.Read(body)
	var user data.User
	json.Unmarshal(body, &user)
	loginSuccess := data.CheckUser(user.Email, user.Password)
	if loginSuccess{
		successMessage := "Login successful!"
		jsData, err := json.Marshal(successMessage)
		if err != nil{
			fmt.Println("Json Marshal Error!")
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsData)
		writer.WriteHeader(200)
	}else{
		errMessage := "Invalid user, try again!"
		jsData, err := json.Marshal(errMessage)
		if err != nil{
			fmt.Println("Json Marshal Error!")
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsData)
		writer.WriteHeader(400)
	}


}


