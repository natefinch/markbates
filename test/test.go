package test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/pkg/browser"
)

func RunTest() {
	box := packr.NewBox("../public")

	http.Handle("/", http.FileServer(box))
	fmt.Println("serving docs at http://localhost:8080")
	fmt.Println("hit ctrl-C to cancel")
	go browser.OpenURL("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
