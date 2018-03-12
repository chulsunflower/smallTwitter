package main

import (
	"net/http"

	"smallTwitter/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/crypto"
	"log"
)

type deletePostTemp struct{
	Email string
	ID uint

}

func post(w http.ResponseWriter, r *http.Request){
	var err error
	switch r.Method {
	case "GET":
		fmt.Println("Using get method!")
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "DELETE":
		fmt.Println("Using delete method!")
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//GET http://localhost:8080/post?email=chulsunflower@163.com
func handleGet(w http.ResponseWriter, r *http.Request)(err error){
	email := r.URL.Query().Get("email")
	userPost, err := data.GetPost(email)
	if err !=nil {
		fmt.Println(err)
	}
	//fmt.Printf("HandleGet %s\n", userPost)
	jsData, err := json.Marshal(userPost)
	if err != nil{
		fmt.Println("Error!")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsData)
	return err

}

//POST http://localhost:8080/post
func handlePost(w http.ResponseWriter, r *http.Request)(err error){
	AccessToken := r.Header.Get("Authorization")
	fmt.Println("HandlePost",AccessToken )
	Valid := ValidateToken(AccessToken)
	if Valid == true{
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		jsonString := make(map[string]string)
		if err := json.Unmarshal(body, &jsonString); err==nil{
			if err := data.SendPost(jsonString["Email"], jsonString["Content"]); err !=nil {
				fmt.Println(err)
			}

		}
		return err
	}
	return
}

//DELETE http://localhost:8080/post
func handleDelete(w http.ResponseWriter, r *http.Request)(err error){
	len := r.ContentLength
	body := make([]byte, len)
	temp := deletePostTemp{}
	r.Body.Read(body)
	if err := json.Unmarshal(body, &temp); err==nil{
		//fmt.Print(temp)
		if err := data.DeletePost(temp.Email,temp.ID); err != nil{
			fmt.Println("Error")
		}

	}else{
		fmt.Println(err)
	}
	return err

}

func ValidateToken(AccessToken string) bool {
	bytes_validate, _ := ioutil.ReadFile("./sample_key.pub")
	rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM(bytes_validate)
	jwt, err := jws.ParseJWT([]byte(AccessToken))
	if err != nil {
		log.Fatal("JWS PARSE ERROR",err)
	}

	// Validate token
	if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
		log.Fatal("Validate Error", err)
	}else{
		fmt.Println("Valid Token!")
		return true
	}
	return false
}

