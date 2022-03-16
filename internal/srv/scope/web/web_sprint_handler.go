package web

import (
	"fmt"
	"net/http"

	aur "prb/internal/srv/scope/common/auto_registerer"

	"github.com/gorilla/mux"
)


type WebSprintHandler struct{
	*aur.AutoRegistereesShared
}

func NewWebSprintHandler(sh *aur.AutoRegistereesShared) *WebSprintHandler {
	return &WebSprintHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *WebSprintHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if sprintNo, ok := vars["sprint"]; ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<h1>%s</h1>", sprintNo)
		return
	}

	Custom404Page(h.AutoRegistereesShared, w, "/prb/p", "projects")
}