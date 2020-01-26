package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"server/api"
)

func init() {

}

func handleResponse(resp api.Response, w http.ResponseWriter) {
	j, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("response: %v", string(j))
	if resp.Success() {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(j)
	w.Write([]byte{'\n'})
}

// ValidationDecoder is a wrapper around a json decoder, so that I can perform request validations
// automatically after decoding. Prevents lots of boiler-plate code
type ValidationDecoder struct {
	*json.Decoder
}

// DecodeAndValidate first calls json.Decode, and then preforms request validation
func (vd ValidationDecoder) DecodeAndValidate(v api.Request) error {
	err := vd.Decode(v)
	if err != nil {
		return err
	}

	return v.Validate()
}

// NewValidationDecoder builds a Validation decoder from a request
func NewValidationDecoder(r *http.Request) ValidationDecoder {
	return ValidationDecoder{json.NewDecoder(r.Body)}
}
