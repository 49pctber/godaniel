package main

import (
	"fmt"
	"net/http"

	"github.com/49pctber/godaniel"
)

func main() {
	http.HandleFunc("/", godaniel.Handler)

	ps := fmt.Sprintf(":%d", godaniel.Port) // port string
	fmt.Printf("Serving on http://localhost%s\n", ps)

	panic(http.ListenAndServe(ps, nil))
}
