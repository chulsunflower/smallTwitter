package main

import(
	"net/http"
)

func main(){
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8080", nil)
}