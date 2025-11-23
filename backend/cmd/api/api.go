package api

import (
	"backend/service/user"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIserver struct {
	addr string
	db *sql.DB
}

func NewAPIserver (addr string , db *sql.DB) *APIserver{
	return &APIserver{
		addr :addr,
		db:  db,
	} //and give me its memory address so I can return a pointer to it.”
}

func (s *APIserver) Run() error {
	router :=mux.NewRouter()
	subrouter := router.PathPrefix("api/v1").Subrouter()  
	
	userHaddler := user.Newhaddler()
	userHaddler.RegisterRoutes(subrouter)

	log.Println("Listening on ",s.addr)

	return http.ListenAndServe(s.addr , router)
	
}