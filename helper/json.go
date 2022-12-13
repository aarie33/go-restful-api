package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, v interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(v)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, v interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(v)
	PanicIfError(err)
}
