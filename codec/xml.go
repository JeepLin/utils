package codec

import (
	"encoding/xml"
)

var XML xmlCodec

type xmlCodec struct {}

func (self xmlCodec) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (self xmlCodec) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}
