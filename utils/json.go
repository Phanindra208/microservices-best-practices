package utils

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/go-openapi/runtime"
	jsoniter "github.com/json-iterator/go"
)

// JSON is a fast replacement for "encoding/json".
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

// NewJSONReader returns a new reader for given value.
func NewJSONReader(v interface{}) io.Reader {
	b, _ := json.Marshal(v)
	return bytes.NewBuffer(b)
}

// JSONConsumer creates a new JSON consumer.
func JSONConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		dec := JSON.NewDecoder(reader)
		dec.UseNumber() // preserve number formats
		return dec.Decode(data)
	})
}

// JSONProducer creates a new JSON producer.
func JSONProducer() runtime.Producer {
	return runtime.ProducerFunc(func(writer io.Writer, data interface{}) error {
		enc := JSON.NewEncoder(writer)
		enc.SetEscapeHTML(false)
		return enc.Encode(data)
	})
}

// JSONTransform transforms type using JSON. Expects two pointer.
func JSONTransform(type1, type2 interface{}) {
	b, _ := JSON.Marshal(type1)
	JSON.Unmarshal(b, type2)
}
