package routes

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Register(){
	router := mux.NewRouter()
	finalHandler := http.HandlerFunc(final)
	router.Handle("/test",middlewareOne(finalHandler))
	http.Handle("/", router)
}

func middlewareOne(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Executing middlewareOne")
    next.ServeHTTP(w, r)
    log.Println("Executing middlewareOne again")
  })
}

func final(w http.ResponseWriter, r *http.Request) {
  log.Println("Executing finalHandler")
  w.Write([]byte("OK"))
}