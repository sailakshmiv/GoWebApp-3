package controllers

import (
	"html/template"
	"net/http"
	"os"
)

type BaseController struct {
	templates      *template.Template
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (this *BaseController) PopulateTemplates(directory string, templates *template.Template) {

	this.populateTemplates(directory, templates)
	this.populateTemplates("templates/shared", templates)
}

func (this *BaseController) populateTemplates(directory string, templates *template.Template) {

	basePath := directory
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}
	templates.ParseFiles(*templatePaths...)
}

func (this *BaseController) View(w http.ResponseWriter, layout, view string, viewModel interface{}) {
	template := this.templates.Lookup(layout)
	template.ParseFiles(view)
	template.Execute(w, viewModel)
}
