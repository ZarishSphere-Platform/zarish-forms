package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zarish-forms <command>")
		return
	}

	switch os.Args[1] {
	case "validate":
		// validate JSON schemas
	case "compile":
		// spreadsheet => schema
	case "preview":
		// generate previews
	default:
		fmt.Println("unknown command")
	}
}