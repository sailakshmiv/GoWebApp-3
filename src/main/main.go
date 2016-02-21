package main

import (
  "net/http"
  "routes"
)

func main() {
	
	routes.Register();
	http.ListenAndServe(":8000", nil)
}