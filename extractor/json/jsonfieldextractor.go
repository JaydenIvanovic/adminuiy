package json

import (
	"io"
	"os/exec"
)

// Probably jumped the gun building this...
// The client will fetch the data, so it
// needs to resolve the jsonPath in javascript
func ExtractJson(jsonBlob []byte, jsonPath string) string {
	cmd := exec.Command("jq", jsonPath)

	input, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	io.WriteString(input, string(jsonBlob))
	input.Close()

	result, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return string(result)
}
