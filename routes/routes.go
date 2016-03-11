package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/yanndr/GoWebApp/controllers"
	"github.com/yanndr/GoWebApp/middleware"
)

func Register() {

	router := mux.NewRouter()

	commonHandlers := alice.New(handlers.RecoveryHandler(), middleware.LoggingHandler, handlers.CompressHandler)
	authHandlers := commonHandlers.Append(middleware.AuthHandler)
	loginHandlers := commonHandlers.Append(middleware.RedirectHandler)
	postLoginHandlers := commonHandlers.Append(middleware.PostLoginRedirectHandler)

	homeController := controllers.NewHomeController()
	accountController := controllers.NewAccountController()

	http.Handle("/lib/", http.StripPrefix("/lib/", http.FileServer(http.Dir("wwwroot/lib"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("wwwroot/img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("wwwroot/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("wwwroot/js"))))

	router.Handle("/", authHandlers.ThenFunc(homeController.Index))
	router.Handle("/Home", authHandlers.ThenFunc(homeController.Index))
	router.Handle("/Account/Login", loginHandlers.ThenFunc(accountController.Login)).Methods("GET")
	router.Handle("/Account/Login", postLoginHandlers.ThenFunc(accountController.PostLogin)).Methods("POST")
	http.Handle("/", router)

}
