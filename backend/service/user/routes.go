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

func (h *haddler ) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.haddlerLogin).Methods("POST")
	router.HandleFunc("/register",h.haddlerLogin).Methods("POST")
	router.HandleFunc("/login",h.haddlerLogin).Methods("POST")

}

func (h *haddler) haddlerLogin(w http.ResponseWriter , r *http.Request){

}

func (h *haddler) haddlerRegister(w http.ResponseWriter , r *http.Request){
	
}
