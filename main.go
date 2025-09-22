package main

import(
	"net/http"
	"log"
)


func main(){

	mu := http.NewServeMux()
	mu.Handle("/", http.FileServer(http.Dir(".")))
	server := &http.Server {
		Addr:		":8080",
		Handler:	mu,
	}
	log.Fatal(server.ListenAndServe())
}
