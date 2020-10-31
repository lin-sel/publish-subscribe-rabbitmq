package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	er "github.com/lin-sel/contact-app/error"
)

// UnmarshalRequestBody Parse Request Body
func UnmarshalRequestBody(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return er.NewHTTPError("request body has empty", http.StatusNoContent)
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return er.NewHTTPError("unable to read request body", http.StatusInternalServerError)
	}

	if len(requestBody) == 0 {
		return er.NewHTTPError("request body has length 0", http.StatusNoContent)
	}

	err = json.Unmarshal(requestBody, target)
	if err != nil {
		return er.NewHTTPError("unable to parse request body", http.StatusInternalServerError)
	}

	return nil
}
