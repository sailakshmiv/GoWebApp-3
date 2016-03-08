package controllers

import (
	"html/template"
	"net/http"

	"github.com/yanndr/GoWebApp/security"
)

type AccountController struct {
	BaseController
}

func NewAccountController() *AccountController {
	controller := new(AccountController)
	controller.templates = template.New("templates")
	controller.PopulateTemplates("templates/account", controller.templates)

	return controller
}

func (this *AccountController) Login(w http.ResponseWriter, r *http.Request) {
	security.GetInstance().CreateCookie(w)
	template := this.templates.Lookup("_layout.html")
	template.ParseFiles("index.html")
	template.Execute(w, nil)

}

func (this *AccountController) PostLogin(w http.ResponseWriter, r *http.Request) {
	template := this.templates.Lookup("_layout.html")
	template.ParseFiles("index.html")
	template.Execute(w, nil)
}
