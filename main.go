package main

import (
	"net/http"

	"github.com/yanndr/GoWebApp/routes"
)

func main() {

	routes.Register()
	http.ListenAndServe(":8000", nil)
}
