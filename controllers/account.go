package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/yanndr/GoWebApp/models"
	"github.com/yanndr/GoWebApp/security"
	"github.com/yanndr/GoWebApp/viewmodels"
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
	vm := new(viewmodels.Login)
	vm.Title = "Login"
	vm.ReturnUrl = r.URL.Query()["returnUrl"][0]
	this.View(w, "_layout.html", "login.html", vm)
}

func (this *AccountController) PostLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if user, err := models.Login(username, password); err == nil {
		log.Print(user)
		security.GetInstance().CreateCookie(w)

		returnUrl := r.URL.Query()["returnUrl"]

		log.Printf("returnUrl:%s", returnUrl)

		/*decoded, err := base64.StdEncoding.DecodeString(returnUrl)
		if err != nil && len(decoded) == 0 {
			http.Redirect(w, r, "/", 302)
		} else {
			log.Printf("decoded:%s", string(decoded))

			http.Redirect(w, r, string(decoded), 302)
		}*/
	}

	this.Login(w, r)
}

/*func redirect(returnUrl string) {
	if returnUrl != "" {
		http.Redirect(w, r, returnUrl, 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}*/
