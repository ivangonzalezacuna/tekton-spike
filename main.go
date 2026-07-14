package main

import (
	"fmt"
	"os"
)

const unused = `unusedstr`

func main() {
	fmt.Fprintf(os.Stdout, "Hello world")
}
