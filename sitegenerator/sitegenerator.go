package sitegenerator

import (
	"fmt"
	"os"

	"github.com/JaydenIvanovic/adminuiy/models"
	"github.com/JaydenIvanovic/adminuiy/templatefs"
)

func BuildStaticFiles(root models.Root) {
	mustCreateOutputDir()
	buildRootPage(root)
	parseElements(root)
}

func parseElements(root models.Root) {
	for _, element := range root.Elements {
		e := Element{
			Name: element.(map[string]interface{})["name"].(string),
			Type: element.(map[string]interface{})["type"].(string),
			Data: element,
		}
		switch e.Type {
		case FormElement:
			createElement(root, e, "form")
		case ListElement:
			createElement(root, e, "list")
		case CrudElement:
			crudViews := []string{"crud_list", "crud_detail"}
			for _, view := range crudViews {
				createElement(root, e, view)
			}
		default:
			panic(fmt.Sprintf("passed element (%s) which doesn't exist", e.Name))
		}
	}
}

func mustCreateOutputDir() {
	err := os.MkdirAll("output", os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func mustCreateHtmlFile(fileName string) *os.File {
	f, err := os.Create(fmt.Sprintf("output/%s.html", fileName))
	if err != nil {
		panic(err)
	}
	return f
}

func buildRootPage(root models.Root) {
	rv := models.RootView{
		Heading:  root.Title,
		Branding: root.Branding,
		Tiles:    []models.Tile{},
	}

	for _, element := range root.Elements {
		rv.Tiles = append(rv.Tiles, models.Tile{
			Url:         fmt.Sprintf("./%s.html", element.(map[string]interface{})["name"]),
			Heading:     element.(map[string]interface{})["heading"].(string),
			Description: element.(map[string]interface{})["description"].(string),
		})
	}

	f := mustCreateHtmlFile("root")
	templateFile := templatefs.GetElementTemplatePath("root")
	tmpl := templatefs.MustParseTemplateFile(templateFile)
	if err := tmpl.Execute(f, rv); err != nil {
		panic(err)
	}
}

func createElement(root models.Root, element Element, templateName string) {
	templateFile := templatefs.GetElementTemplatePath(templateName)
	tmpl := templatefs.MustParseTemplateFile(templateFile)

	file := mustCreateHtmlFile(element.Name)
	data := struct {
		Branding models.Branding
		Data     interface{}
	}{
		Branding: root.Branding,
		Data:     element.Data,
	}

	if err := tmpl.Execute(file, data); err != nil {
		panic(err)
	}
}
