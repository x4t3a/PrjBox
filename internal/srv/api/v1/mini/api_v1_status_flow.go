package apiv1mini

import (
	"encoding/json"
	"net/http"
	"prb/internal/srv/api/v1/fields"
	aur "prb/internal/srv/common/auto_registerer"
	"prb/internal/srv/common/stub"
)

type APIV1MiniStatusFlowHandler struct {
	*aur.AutoRegistereesShared
}

func NewAPIV1MiniStatusFlowHandler(sh *aur.AutoRegistereesShared) *APIV1MiniStatusFlowHandler {
	return &APIV1MiniStatusFlowHandler{
		AutoRegistereesShared: sh,
	}
}

var flows = map[string]map[string]string{
	fields.FieldStatusOpen.String():       {fields.FieldStatusInProgress.String(): "all", fields.FieldStatusClosed.String(): "all"},
	fields.FieldStatusInProgress.String(): {fields.FieldStatusOpen.String(): "all", fields.FieldStatusClosed.String(): "all"},
	fields.FieldStatusClosed.String():     {fields.FieldStatusReOpened.String(): "all"},
}

func (h *APIV1MiniStatusFlowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Status string
	}{}

	if stub.UserHasRights() {
		switch r.Method {
		case "GET":
			if vals, ok := r.URL.Query()["status"]; ok && len(vals) > 0 {
				req.Status = vals[0]
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(flows[req.Status])
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
