package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate renders templates using html
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache(w)
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser : ", err )
	}
}
// CreateTemplateCache is fancy way of rendering templates 
func CreateTemplateCache(w http.ResponseWriter) (map[string]*template.Template, error) {
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

