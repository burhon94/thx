package app

import (
	"errors"
	"fmt"
	"github.com/burhon94/thx/pkg/rest"
	"github.com/burhon94/thx/pkg/structs"
	"log"
	"net/http"
)

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	var resp structs.ResponseAPI

	resp.Code = 200
	resp.Payload = "Hello, from THX :)"
	resp.Info = "Hello, from THX :)"
	resp.Errors = errors.New("hello, from THX :)")

	err := rest.WriteBody(w, 202, resp)
	if err != nil {
		log.Printf("can't response: %v", err)
	}
}

func (routers *routes) handleTest(w http.ResponseWriter, _ *http.Request) {

	var resp structs.ResponseAPI

	err := routers.DBService.add()
	if err != nil {
		resp.Code = 500
		resp.Errors = err

		err := rest.WriteBody(w, 500, resp)
		if err != nil {
			log.Printf("can't response: %v", err)
		}

		return
	}

	resp.Code = 200
	resp.Payload = "add Success"

	err = rest.WriteBody(w, 200, resp)
	if err != nil {
		log.Printf(" can't response: %v", resp)
	}
}

func (routers *routes) SayThx(w http.ResponseWriter, r *http.Request) {
	var (
		req  structs.ReqData
		resp structs.ResponseAPI
	)

	err := rest.ReadBody(r, req)
	if err != nil {
		log.Printf("")
		resp.Code = 400
		resp.Errors = errors.New("error.BadRequest")

		err := rest.WriteBody(w, 400, resp)
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}

		return
	}

	err = routers.DBService.sayThx()
	if err != nil {
		resp.Code = 500
		resp.Errors = errors.New(fmt.Sprintf("error.SayThx: %s", err.Error()))

		err = rest.WriteBody(w, 500, resp)
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}

		return
	}

	resp.Code = 202
	resp.Payload = "OK, Success"
	resp.Info = "say thx :)"

	err = rest.WriteBody(w, 202, resp)
	if err != nil {
		log.Printf("can't sent response: %v", err)
	}
}
