package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alvintoh/go-programming-projects/simple-http-server-go/internal/app/handler"
)

func main() {
	fileServer := http.FileServer(http.Dir("./internal/web/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handler.FormHandler)
	http.HandleFunc("/hello", handler.HelloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
