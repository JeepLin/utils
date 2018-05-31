package http

import (
	"net/http"
	"github.com/JeepLin/utils/slice"
)

type HttpHandleFunc http.HandlerFunc

func WithIPLimit(iplist []string, f HttpHandleFunc) HttpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request){
		found := slice.NewStringSlice(iplist).Contains(GetIPAddr(r))
		if !found {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		f(w, r)
		return
	}
}