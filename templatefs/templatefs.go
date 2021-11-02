package templatefs

import (
	"fmt"
	"html/template"
	"path"
)

func GetElementTemplatePath(elementName string) string {
	return fmt.Sprintf("templates/elements/%s.tmpl.html", elementName)
}

func MustParseTemplateFile(fileName string) *template.Template {
	tmpl, err := template.New(path.Base(fileName)).ParseGlob("templates/partials/*")
	if err != nil {
		panic(err)
	}
	tmpl, err = tmpl.ParseFiles(fileName)
	if err != nil {
		panic(err)
	}
	return tmpl
}
