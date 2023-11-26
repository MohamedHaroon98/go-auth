package routes

import (
	"github.com/MohamedHaroon98/go-auth/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.RegisterAccount).Methods("POST")
	r.HandleFunc("/login", controllers.LoginAccount).Methods("POST")
	r.HandleFunc("/listusers", controllers.GetAllUsernames).Methods("GET")

	return r
}
