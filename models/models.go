package models

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
