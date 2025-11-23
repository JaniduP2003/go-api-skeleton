package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type haddler struct{

}

func Newhaddler (     )*haddler{
	return &haddler{}
}

func (h*Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.handleLogin).Methods("POST")
	router.HandleFunc("/register",h.handleLogin).Methods("POST")
	router.HandleFunc("/login",h.handleLogin).Methods("POST")

}

func (h *haddler) haddlerLogin(w *http.ResponseWriter , r *http.Request){
	
}