package routes

import (
	"github.com/MohamedHaroon98/go-auth/controllers"
	"github.com/MohamedHaroon98/go-auth/utils"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.RegisterAccount).Methods("POST")
	r.HandleFunc("/login", controllers.LoginAccount).Methods("POST")
	r.HandleFunc("/listusers", controllers.GetAllUsernames).Methods("GET")
	r.HandleFunc("/health", utils.HealthCheckDatabase).Methods("GET")

	return r
}
