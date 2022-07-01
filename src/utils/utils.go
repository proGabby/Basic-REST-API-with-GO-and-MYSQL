package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//parse the body
func ParseBody(r *http.Request, x any) {
	//read from r until an error or EOF and returns the data it read
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		//Unmarshal parses the JSON-encoded data
		//and stores the result in the value pointed to by x.
		//If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
