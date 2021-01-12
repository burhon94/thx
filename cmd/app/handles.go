package app

import "net/http"

func handleHealth(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("it work"))
	w.WriteHeader(202)
}
