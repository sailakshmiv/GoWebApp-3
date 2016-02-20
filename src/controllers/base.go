package controllers

import (
  "os"
  "html/template"
)

type BaseController struct {
	 *template.Template
	}

func (this *BaseController) PopulateTemplates(directory string, templates *template.Template,templatePaths *[]string) {

	basePath := directory
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	
	templatePathsRaw, _ := templateFolder.Readdir(-1)
	
	
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			//fmt.Println(pathInfo.Name())
			*templatePaths = append(*templatePaths,basePath + "/" + pathInfo.Name())
		}/*else{
			populateTemplates(basePath + "/" + pathInfo.Name(),templates,templatePaths)
		}*/
	}
	
	templates.ParseFiles(*templatePaths...)
	
}


