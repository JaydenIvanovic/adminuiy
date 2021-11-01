package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

type Root struct {
	Title    string
	Branding Branding
	Elements []interface{}
}

type Form struct {
	Name        string
	Heading     string
	Description string
	Fields      []Field
	Endpoint    string
}

type FormView struct {
	Branding
	Form
}

type RootView struct {
	Branding
	Heading string
	Tiles   []Tile
}

type Tile struct {
	Url         string
	Heading     string
	Description string
}

type Field struct {
	Label    string
	Name     string
	Type     string
	HelpText string
}

type Branding struct {
	LogoUrl     string
	LogoWidth   string
	LogoHeight  string
	CompanyName string
}

func main() {
	uiFile := os.Args[1]
	root := createRootStructFromFile(uiFile)
	buildStaticFiles(root)
}

func createRootStructFromFile(uiFilePath string) Root {
	rawFileContents, err := os.ReadFile(uiFilePath)
	if err != nil {
		panic(err)
	}

	var root Root
	err = json.Unmarshal(rawFileContents, &root)
	if err != nil {
		panic(err)
	}

	return root
}

func buildStaticFiles(root Root) {
	buildRootPage(root)
	parseElements(root)
}

func buildRootPage(root Root) {
	mustCreateOutputDir()
	f := mustCreateHtmlFile("root")
	rv := RootView{
		Heading:  root.Title,
		Branding: root.Branding,
		Tiles:    []Tile{},
	}
	for _, element := range root.Elements {
		rv.Tiles = append(rv.Tiles, Tile{
			Url:         fmt.Sprintf("./%s.html", element.(map[string]interface{})["name"]),
			Heading:     element.(map[string]interface{})["heading"].(string),
			Description: element.(map[string]interface{})["description"].(string),
		})
	}

	templateFile := "templates/root.tmpl.html"
	rawTeplateFileContents, err := os.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("root").Parse(string(rawTeplateFileContents))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(f, rv); err != nil {
		panic(err)
	}
}

func parseElements(root Root) {
	for _, element := range root.Elements {
		switch element.(map[string]interface{})["type"].(string) {
		case "form":
			var form Form
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

func createForm(root Root, form Form) {
	templateFile := "templates/form.tmpl.html"
	rawTeplateFileContents, err := os.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("ui").Parse(string(rawTeplateFileContents))
	if err != nil {
		panic(err)
	}

	mustCreateOutputDir()

	f := mustCreateHtmlFile(form.Name)
	fv := FormView{
		Branding: root.Branding,
		Form:     form,
	}
	if err = tmpl.Execute(f, fv); err != nil {
		panic(err)
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
