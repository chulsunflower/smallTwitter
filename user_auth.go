package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"smallTwitter/data"
	"io/ioutil"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/crypto"
	"time"
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
		responseMap := make(map[string]string)
		responseMap["SuccessMessage"] =  "Login successful!"
		AccessToken := generateToken()
		responseMap["AccessToken"] = AccessToken
		jsData, err := json.Marshal(responseMap)
		if err != nil{
			fmt.Println("Json Marshal Error!")
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsData)
		writer.WriteHeader(200)
		fmt.Println(AccessToken)
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

func generateToken()string {
	bytes_generate, _ := ioutil.ReadFile("./sample_key.priv")

	claims := jws.Claims{}
	claims.SetExpiration(time.Now().Add(time.Duration(600) * time.Second))

	rsaPrivate, _ := crypto.ParseRSAPrivateKeyFromPEM(bytes_generate)
	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	b, _ := jwt.Serialize(rsaPrivate)
	return string(b)

}
