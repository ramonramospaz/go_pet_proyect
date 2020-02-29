package util

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

// Store files in gob format
func Store(data interface{}, filename string) (err error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err = encoder.Encode(data)
	if err == nil {
		err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	}
	return
}

// Load files in gob format
func Load(data interface{}, filename string) (err error) {
	raw, err := ioutil.ReadFile(filename)
	if err == nil {
		buffer := bytes.NewBuffer(raw)
		dec := gob.NewDecoder(buffer)
		err = dec.Decode(data)
	}
	return
}
