package gkv

import (
	"net/http"
)

type Master struct {
}

func (m *Master) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check request url path then route to handler

}
