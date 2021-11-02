package sitegenerator

import (
	"encoding/json"
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

func createForm(root models.Root, form models.Form) {
	templateFile := templatefs.GetElementTemplatePath("form")
	tmpl := templatefs.MustParseTemplateFile(templateFile)

	f := mustCreateHtmlFile(form.Name)
	fv := models.FormView{
		Branding: root.Branding,
		Form:     form,
	}
	if err := tmpl.Execute(f, fv); err != nil {
		panic(err)
	}
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

func parseElements(root models.Root) {
	for _, element := range root.Elements {
		switch element.(map[string]interface{})["type"].(string) {
		case "form":
			var form models.Form
			rawData, err := json.Marshal(element)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(rawData, &form)
			if err != nil {
				panic(err)
			}
			createForm(root, form)
		default:
			panic("passed element which doesn't exist")
		}
	}
}
