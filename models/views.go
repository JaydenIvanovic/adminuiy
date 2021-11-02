package models

type FormView struct {
	Branding
	Form
}

type RootView struct {
	Branding
	Heading string
	Tiles   []Tile
}
