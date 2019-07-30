package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// SendError : SendError
func SendError(w http.ResponseWriter, status int, err interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

// SendSuccess : SendSuccess
func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

//LogFatal : LogFatal
// If for any reason, something goes amiss then exit and
// send the message to the logging console.
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
