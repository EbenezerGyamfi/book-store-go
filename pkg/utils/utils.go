package utils

import (
	"encoding/json"
	"io"
	"net/http"
)


func ParseBody(r *http.Request, x interface{})  {

	body, err := io.ReadAll(r.Body)

	if err!= nil {
        panic(err)
    }
	defer r.Body.Close()

	err = json.Unmarshal(body, x)

	if err!= nil {
        panic(err)
    }

}