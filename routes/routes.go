package routes

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/yanndr/GoWebApp/controllers"
	"github.com/yanndr/GoWebApp/middleware"
)

func Register() {

	commonHandlers := alice.New(middleware.LoggingHandler, middleware.RecoverHandler)
	authHandlers := commonHandlers.Append(middleware.AuthHandler)

	homeController := controllers.NewHomeController()
	accountController := controllers.NewAccountController()

	router := mux.NewRouter()
	router.Handle("/", authHandlers.ThenFunc(homeController.Index))

	router.Handle("/Account/Login", commonHandlers.ThenFunc(accountController.Login)).Methods("GET")
	router.Handle("/Account/Login", commonHandlers.ThenFunc(accountController.PostLogin)).Methods("POST")

	http.Handle("/", router)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/lib/", serveResource)

}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "wwwroot" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
