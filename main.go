package httpservergo

import(
	"net/http"
	"log"
)

func main(){
	mu := http.NewServeMux()

	server := &http.Server {
		Addr:		":8080",
		Handler:	mu,
	}
	log.Fatal(server.ListenAndServe())
}
