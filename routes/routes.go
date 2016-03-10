package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/yanndr/GoWebApp/controllers"
	"github.com/yanndr/GoWebApp/middleware"
)

func Register() {

	commonHandlers := alice.New(middleware.LoggingHandler, middleware.RecoverHandler, middleware.Gzip)
	authHandlers := commonHandlers.Append(middleware.AuthHandler)

	homeController := controllers.NewHomeController()
	accountController := controllers.NewAccountController()

	http.Handle("/lib/", http.StripPrefix("/lib/", http.FileServer(http.Dir("wwwroot/lib"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("wwwroot/img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("wwwroot/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("wwwroot/js"))))

	router := mux.NewRouter()
	router.Handle("/", authHandlers.ThenFunc(homeController.Index))
	router.Handle("/Account/Login", commonHandlers.ThenFunc(accountController.Login)).Methods("GET")
	router.Handle("/Account/Login", commonHandlers.ThenFunc(accountController.PostLogin)).Methods("POST")
	http.Handle("/", router)

}
