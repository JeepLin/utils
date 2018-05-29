package codec

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal(data []byte, v interface{}) error
}