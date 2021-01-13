package rest

import (
	"encoding/json"
	"github.com/burhon94/thx/pkg/structs"
	"io/ioutil"
	"net/http"
)

func ReadBody(r *http.Request, value interface{}) (err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &value)
	if err != nil {
		return err
	}

	return nil
}

func WriteBody(w http.ResponseWriter, code int, body structs.ResponseAPI) (err error) {

	marshal, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = w.Write(marshal)
	w.WriteHeader(code)

	return nil

}
