package scheme

import (
	"encoding/json"
	"reflect"
)

type Request struct {
	Method string            `json:"method"`
	Args   []json.RawMessage `json:"args"`
}

type Response struct {
	Status int    `json:"status"`
	Value  any    `json:"value"`
	Error  string `json:"error,omitempty"`
}

type Method struct {
	Name    string
	ArgType reflect.Type
	Args    []reflect.Value
	Count   int
}
