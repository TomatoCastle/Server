//ファイルだけです
package main

import (
	"net/http"

	"./handlers"
)


func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/callback", handlers.Callback)
	server.ListenAndServe()
	initSesion()
}
