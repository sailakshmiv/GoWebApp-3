package controllers

import (
	"html/template"
	"os"
)

type BaseController struct {
	templates *template.Template
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
		} /*else{
			populateTemplates(basePath + "/" + pathInfo.Name(),templates,templatePaths)
		}*/
	}
	templates.ParseFiles(*templatePaths...)
}
