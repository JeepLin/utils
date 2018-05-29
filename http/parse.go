package http

import (
	"fmt"
	"net/http"
	"errors"
	"io/ioutil"
	"utils/codec"
)

//ParseAndUnmarshal 解析并序列化请求
func ParseAndUnmarshal(r *http.Request, arg interface{}, fn codec.Unmarshaler) (b []byte, err error) {
	if r == nil {
		err = errors.New("[ParseAndUnmarshal] http.Request is nil")
		return
	}

	if fn == nil {
		err = errors.New("[ParseAndUnmarshal] Unmarshaler is nil")
		return
	}

	if r.Body == nil {
		err = errors.New("[ParseAndUnmarshal] body is nil")
		return
	}

	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		r.Body.Close()
		err = fmt.Errorf("[ParseAndUnmarshal] ReadAll err:%s", err.Error())
		return
	}

	r.Body.Close()

	err = fn.Unmarshal(b, &arg)
	if err != nil {
		err = fmt.Errorf("[ParseAndUnmarshal] Unmarshal err:%s, body:%s",
						err.Error(), string(b))
		return
	}

	return
}
