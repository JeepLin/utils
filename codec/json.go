package codec

import "encoding/json"

var JSON jsonCodec

type jsonCodec struct {}

func (self jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (self jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}