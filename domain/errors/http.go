package errors

type HttpErrorObject struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type HttpError struct {
	Error HttpErrorObject `json:"error"`
}
