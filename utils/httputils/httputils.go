package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sshindanai/bookstore-utils-go/resterrors"
)

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err resterrors.RestErr) {
	fmt.Println(err.StatusCode(), err.Causes())
	RespondJSON(w, err.StatusCode(), err)
}
