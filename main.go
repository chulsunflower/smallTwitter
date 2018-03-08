package main

import(
	"net/http"
)

func main(){
	http.HandleFunc("/login", login)
	//http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)
	//http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/post", post)

	http.ListenAndServe(":8080", nil)
}