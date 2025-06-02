package services

import (
	"bytes"
	"encoding/json"
	"training-backend/package/log"
)

func Decode(in, out interface{}) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		log.Errorf("error creating new encoder: %v", err)
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		log.Errorf("error creating new decoder: %v", err)
	}
}
