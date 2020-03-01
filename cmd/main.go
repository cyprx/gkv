package main

import (
	"net/http"

	"github.com/cyprx/gkv"
)

func main() {

	m := &gkv.Master{}
	http.ListenAndServe(":7000", m)

}
