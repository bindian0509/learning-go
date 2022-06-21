package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bindian0509/learning-go/web-app-html/pkg/config"
	"github.com/bindian0509/learning-go/web-app-html/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplate sets the config for the template package 
func NewTemplate(a *config.AppConfig) {
	app = a 
}

func AddDefaultData (td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser : ", err )
	}
}
// CreateTemplateCache is fancy way of rendering templates 
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.htm")
	if err != nil {
		return myCache, err 
	}
	for _, page := range pages {
		//fmt.Println(":: Now Rendering page : ", page)
		name := filepath.Base(page)
		tset, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err 
		}

		matches, err := filepath.Glob("./templates/*.layout.htm")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			tset, err  = tset.ParseGlob("./templates/*.layout.htm")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = tset 
	}
	return myCache, nil
}

