package authtest

import (
	"log"
	"net/http"

	"github.com/kanatatsu64/refrector/auth"
)

// InitAuthTestServer is a test mok of server.InitServer.
func InitAuthTestServer(port string) {
	auth.InitAuth()

	http.HandleFunc("/home", GoogleAuthTestHandler)

	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
