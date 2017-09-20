package main

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("public/index.html")
	fmt.Println("is embedded: ", box.Has("public/index.html"))
}
