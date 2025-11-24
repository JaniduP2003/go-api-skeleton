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
	router.HandleFunc("/register",h.haddlerRegister).Methods("POST")

}

func (h *haddler) haddlerLogin(w http.ResponseWriter , r *http.Request){

}

func (h *haddler) haddlerRegister(w http.ResponseWriter , r *http.Request){
	//get json payload (for this use type make a type folder and add the code)
	//ceack if the user exsits
	 //if it dose not the create him
	
}
