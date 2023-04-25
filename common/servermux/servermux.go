package servermux

import "net/http"

var appServerMux *http.ServeMux

func AppServerMux() *http.ServeMux {
	if appServerMux != nil {
		return appServerMux
	}
	appServerMux = http.NewServeMux()
	return appServerMux
}
