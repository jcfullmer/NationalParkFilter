package main

import "net/http"

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the root."))
}
