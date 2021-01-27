package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/sshindanai/bookstore-utils-go/resterrors"
)

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err *resterrors.RestErr) {
	RespondJSON(w, err.Code, err)
}
