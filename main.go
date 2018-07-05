package main

import (
	"fmt"
	"os"
)

const (
	// DefaultVersion is a version which is used by this package
	// for all the types of sent/received data.
	DefaultVersion = "2.0"
)

func main() {
	fmt.Fprintf(os.Stdout, "Exected response version to be %d but it is %d", DefaultVersion, DefaultVersion)
}
