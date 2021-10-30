package main

import (
	"encoding/json"
	"html/template"
	"os"
)

type Form struct {
	Branding    Branding
	Heading     string
	Description string
	Fields      []Field
	Endpoint    string
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
	templateFile := os.Args[2]

	rawFileContents, err := os.ReadFile(uiFile)
	if err != nil {
		panic(err)
	}

	var f Form
	err = json.Unmarshal(rawFileContents, &f)
	if err != nil {
		panic(err)
	}

	rawTeplateFileContents, err := os.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("ui").Parse(string(rawTeplateFileContents))
	if err != nil {
		panic(err)
	}

	if err = tmpl.Execute(os.Stdout, f); err != nil {
		panic(err)
	}
}
