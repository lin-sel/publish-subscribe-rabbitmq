package util

import (
	"encoding/json"
	"net/http"

	err "github.com/lin-sel/contact-app/error"
)

// RespondJSON Marshal Respond data & write to ResponseWriter
func RespondJSON(w http.ResponseWriter, status int, respond interface{}) {
	resp, err := json.Marshal(respond)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an error occur"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(resp))
}

// RespondErrorMessage makes the error response with payload as json format
func RespondErrorMessage(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

// RespondError returns a validation error else
func RespondError(w http.ResponseWriter, er error) {
	switch er.(type) {
	case err.ValidationError:
		RespondJSON(w, http.StatusBadRequest, er)
	case err.HTTPError:
		httpError := er.(err.HTTPError)
		RespondErrorMessage(w, httpError.HTTPStatus, httpError.HTTPError)
	default:
		RespondErrorMessage(w, http.StatusInternalServerError, "an error occur")
	}
}
