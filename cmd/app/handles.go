package app

import (
	"context"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("it work"))
	w.WriteHeader(202)
}

func (r *routes) handleTest(w http.ResponseWriter, request *http.Request)  {
	sql := `insert into users values (2);`
	_, err := r.pool.Exec(context.Background(), sql)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	w.Write([]byte("add Success"))
	w.WriteHeader(200)
}
