package models

type FormView struct {
	Branding
	Form
}

type ListView struct {
	Branding
	List
}

type RootView struct {
	Branding
	Heading string
	Tiles   []Tile
}
