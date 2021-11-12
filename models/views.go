package models

type FormView struct {
	Branding
	Form
}

type ListView struct {
	Branding
	List
}

type CrudModelView struct {
	Branding
	CrudModel
}

type RootView struct {
	Branding
	Heading string
	Tiles   []Tile
}
