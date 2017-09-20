package main

import (
	"fmt"

	"github.com/gobuffalo/packr"
	"github.com/natefinch/markbates/test"
)

func main() {
	box := packr.NewBox("public/index.html")
	fmt.Println("is embedded: ", box.Has("public/index.html"))
	test.RunTest()
}
