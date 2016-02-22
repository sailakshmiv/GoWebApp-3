package controllers

import (
	"log"
	"net/http"
	"html/template"
)

type HomeController struct {
	BaseController
}

func NewHomeController()(*HomeController){
	controller := new(HomeController)
	controller.templates = template.New("templates")
	controller.PopulateTemplates("templates/home",controller.templates)

	
	return controller;
}

func (this *HomeController) Index(w http.ResponseWriter, r *http.Request){	 
	 template := this.templates.Lookup("test.html")
	 log.Println("template ok")
	 template.Execute(w, nil)
	 
	//w.Write([]byte("OK"))
}