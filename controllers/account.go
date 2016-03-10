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

type LoginViewModel struct {
	Title     string
	ReturnUrl string
}

func (this *AccountController) Login(w http.ResponseWriter, r *http.Request) {
	vm := new(LoginViewModel)
	vm.Title = "Login"

	returnUrl := r.URL.Query()["returnUrl"][0]

	vm.ReturnUrl = returnUrl

	log.Printf("returnUrl:%s", returnUrl)

	template := this.templates.Lookup("_layout.html")
	template.ParseFiles("login.html")
	template.Execute(w, vm)

}

func (this *AccountController) PostLogin(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	if user, err := login(username, password); err == nil {
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

	if username == "yann" && password == "password" {
		result := User{Id: 1, Name: "Yann"}
		return &result, nil
	} else {
		log.Println("Login Failed")
		return nil, errors.New("Nope")
	}

}
