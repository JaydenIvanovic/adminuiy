package sitegenerator

type Element struct {
	Name string
	Type string
	Data interface{}
}

const (
	FormElement = "form"
	ListElement = "list"
	CrudElement = "crud"
)
