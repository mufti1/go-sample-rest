package responsewriter

import (
	"encoding/json"
	"net/http"
	"strings"
)

// ResponseFormat stands for our default response in API
type ResponseFormat struct {
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

// ResponseOK is function to return json OK
func (rf *ResponseFormat) ResponseOK(code int, data interface{}, w http.ResponseWriter) {
	rf.Data = data

	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(rf)
	if err != nil {
		resErr := ResponseFormat{
			Errors: err.Error(),
		}
		jsonErr, _ := json.Marshal(resErr)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonErr)
	}

	w.WriteHeader(code)
	w.Write(resp)
}

// ResponseNOK is function to return json NOK
func (rf *ResponseFormat) ResponseNOK(code int, errors interface{}, w http.ResponseWriter) {
	rf.Errors = errors

	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(rf)
	if err != nil {
		resErr := ResponseFormat{
			Errors: err.Error(),
		}
		jsonErr, _ := json.Marshal(resErr)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonErr)
	}

	w.WriteHeader(code)
	w.Write(resp)
}

// GetLastPath of url
func GetLastPath(path string) string {
	parsed := strings.Split(path, "/")
	return parsed[len(parsed)-1]
}
