// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/MohamedHaroon98/go-auth/routes"
	"github.com/gorilla/handlers"
)

func main() {
	r := routes.SetupRouter()

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", handlers.CORS()(r))
}
