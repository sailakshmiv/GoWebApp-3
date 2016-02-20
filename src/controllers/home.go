package controllers

import (
	"net/http"
)

type HomeController struct {
	
}

func (this *HomeController) Index(w http.ResponseWriter, r *http.Request){
	 w.Write([]byte("OK"))
}