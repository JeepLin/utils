package http

import (
	"net/http"
	"github.com/JeepLin/utils/string"
)

type HttpHandleFunc http.HandlerFunc

func WithIPLimit(iplist []string, f HttpHandleFunc) HttpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request){
		found := string.New(iplist).Contains(GetIPAddr(r))
		if !found {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		f(w, r)
		return
	}
}