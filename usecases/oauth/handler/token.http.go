package oauthhandler

import (
	"net/http"
)

func (handler handler) TokenHttp(response http.ResponseWriter, request *http.Request) {
	err := handler.injections.OauthSrv.HandleTokenRequest(response, request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
