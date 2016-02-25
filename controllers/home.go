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

type homeViewModel struct{
	Title string
}

func (this *HomeController) Index(w http.ResponseWriter, r *http.Request){	 
	 template := this.templates.Lookup("_layout.html")
	 log.Println("template ok")
	 
	 var hvm  homeViewModel
	 hvm.Title ="Home"
	 template.ParseFiles("index.html");
	 template.Execute(w, hvm)
	 
	//w.Write([]byte("OK"))
}