package pgsegment

import (
	"io"

	"github.com/go-pg/pg/v11/pgjson"
	"github.com/segmentio/encoding/json"
)

var _ pgjson.Provider = (*JSONProvider)(nil)

type JSONProvider struct{}

func (JSONProvider) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (JSONProvider) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (JSONProvider) NewEncoder(w io.Writer) pgjson.Encoder {
	return json.NewEncoder(w)
}

func (JSONProvider) NewDecoder(r io.Reader) pgjson.Decoder {
	return json.NewDecoder(r)
}
