package controllers

import (
	"errors"
	"html/template"
	"log"
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

	template := this.templates.Lookup("_layout.html")
	template.ParseFiles("login.html")
	template.Execute(w, nil)

}

func (this *AccountController) PostLogin(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	if user, err := login(username, password); err == nil {
		log.Println(user)
		security.GetInstance().CreateCookie(w)
		http.Redirect(w, r, "/", 302)
	}

	template := this.templates.Lookup("_layout.html")
	template.ParseFiles("login.html")
	template.Execute(w, nil)
}

type User struct {
	Id   int
	Name string
}

func login(username string, password string) (*User, error) {

	//hashedPassword := sha512.Sum512([]byte(password))
	//b64Pass := base64.StdEncoding.EncodeToString(hashedPassword[:])
	log.Println(username)
	log.Println(password)

	if username == "yann" && password == "password" {
		log.Println("Login Ok")
		result := User{Id: 1, Name: "Yann"}
		return &result, nil
	} else {
		log.Println("Login Failed")
		return nil, errors.New("Nope")
	}

}
