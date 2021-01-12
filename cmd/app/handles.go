package app

import (
	"context"
	"encoding/json"
	"io/ioutil"
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

type reqData struct {
	Client1 client `json:"client_1"`
	Client2 client `json:"client_2"`
	Action int `json:"action"`
}

type client struct {
	Id int `json:"id"`
	Name string `json:"name"`
	SurName string `json:"sur_name"`
	CountThx int `json:"count_thx"`
	CountStar int `json:"count_star"`
	CanDo bool `json:"can_do"`
}

func (r *routes) SayThx(w http.ResponseWriter, request *http.Request)  {
	var (
		req reqData
	)

	body, _ := ioutil.ReadAll(request.Body)

	_ = json.Unmarshal(body, &req)

	sql := `update users
set count_star = count_star + 1
where id = 1;`

	_, err := r.pool.Exec(context.Background(), sql)
	if err != nil {
		w.Write([]byte("500"))
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(202)
	w.Write([]byte("OK set"))
}