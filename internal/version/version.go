package version

import (
	"fmt"
	"os"
)

// read version 
func ReadVersion() string {
	content, err := os.ReadFile("assets/version.txt")
	if err != nil {
		fmt.Println("Error reading version file: ", err)
		return string("error")
	}
	return string(content)
}