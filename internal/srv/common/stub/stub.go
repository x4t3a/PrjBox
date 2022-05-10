package stub

import (
	"fmt"
	"net/http"
)

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "NOT IMPLEMENTED")
}

func stubTrue() bool {
	return true
}

func UserHasRights() bool {
	return stubTrue()
}