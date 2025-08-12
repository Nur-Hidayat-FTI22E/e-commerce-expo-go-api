package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) error {
	var writer http.ResponseWriter
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		WriteJSONError(writer, http.StatusOK, "failed to encode JSON response")
	}

	return err
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		WriteJSONError(writer, http.StatusOK, "failed to encode JSON response")
	}
}
