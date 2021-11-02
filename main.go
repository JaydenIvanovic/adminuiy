package main

import (
	"encoding/json"
	"os"

	"github.com/JaydenIvanovic/adminuiy/models"
	"github.com/JaydenIvanovic/adminuiy/sitegenerator"
)

func main() {
	uiFile := os.Args[1]
	root := createRootStructFromFile(uiFile)
	sitegenerator.BuildStaticFiles(root)
}

func createRootStructFromFile(uiFilePath string) models.Root {
	rawFileContents, err := os.ReadFile(uiFilePath)
	if err != nil {
		panic(err)
	}

	var root models.Root
	err = json.Unmarshal(rawFileContents, &root)
	if err != nil {
		panic(err)
	}

	return root
}
