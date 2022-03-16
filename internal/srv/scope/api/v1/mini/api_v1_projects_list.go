package apiv1mini

import (
	"encoding/json"
	"net/http"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/types"
)

type APIV1MiniProjectsListHandler struct {
	*aur.AutoRegistereesShared
}

func NewAPIV1IssueCreateHandler(sh *aur.AutoRegistereesShared) *APIV1MiniProjectsListHandler {
	return &APIV1MiniProjectsListHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *APIV1MiniProjectsListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Projects []types.DBProject `json:"projects"`
	}{
		Projects: make([]types.DBProject, 0, 10),
	}
	if err := h.DB.Select(&data.Projects, "SELECT name, link FROM common.projects ORDER BY link ASC"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
