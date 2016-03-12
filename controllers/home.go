package controllers

import (
	"html/template"
	"net/http"

	"github.com/yanndr/GoWebApp/viewmodels"
)

type HomeController struct {
	BaseController
}

func NewHomeController() *HomeController {
	controller := new(HomeController)
	controller.templates = template.New("templates")
	controller.PopulateTemplates("templates/home", controller.templates)

	return controller
}

func (this *HomeController) Index(w http.ResponseWriter, r *http.Request) {

	var hvm viewmodels.Home
	hvm.Title = "Home"

	this.View(w, "_layout.html", "index.html", hvm)
}
