package oauthhandler

import (
	"net/http"
)

func (handler handler) AuthorizeHttp(response http.ResponseWriter, request *http.Request) {
	err := handler.injections.OauthSrv.HandleAuthorizeRequest(response, request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}
}
