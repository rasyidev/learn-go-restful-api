package helper

import (
	"encoding/json"
	"net/http"
)

// Decode json dari writer
func ReadFromRequestBody(request *http.Request, result interface{}) {
	// menangkap data dari io.writer, decode
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

// Encode json ke dalam writer
func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	// menambahkan header response Content-Type
	writer.Header().Add("Content-Type", "application/json")

	// encode data ke json, kirim ke writer
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
