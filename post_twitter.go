package main

import (
	"net/http"

	"smallTwitter/data"
	"encoding/json"
	"fmt"
)

func post(w http.ResponseWriter, r *http.Request){
	var err error
	switch r.Method {
	case "GET":
		fmt.Println("Using get method!")
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	//case "PUT":
	//	err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request)(err error){
	email := r.URL.Query().Get("email")
	if err := data.GetPost(email); err !=nil {
		fmt.Println(err)
	}

	return err

}

func handlePost(w http.ResponseWriter, r *http.Request)(err error){
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

func handleDelete(w http.ResponseWriter, r *http.Request)(err error){
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	jsonString := make(map[string]string)
	if err := json.Unmarshal(body, &jsonString); err==nil{
		if err := data.DeletePost(jsonString["Email"], jsonString["Content"]); err !=nil {
			fmt.Println(err)
		}

	}
	return err

}

