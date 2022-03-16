package apiv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prb/internal/srv/scope/api/v1/fields"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/stub"
	"prb/internal/srv/scope/common/types"
	"strconv"
	"strings"
)

type APIV1IssueGetHandler struct {
	*aur.AutoRegistereesShared
}

func NewAPIV1IssueGetHandler(sh *aur.AutoRegistereesShared) *APIV1IssueGetHandler {
	return &APIV1IssueGetHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *APIV1IssueGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := struct {
		types.DBProjectLink
		types.DBIssueID
	}{}

	if stub.UserHasRights() {
		switch r.Method {
		case "GET":
			if vals, ok := r.URL.Query()[req.DBProjectLink.GetMetaName()]; ok && len(vals) == 1 {
				req.Link = vals[0]
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if vals, ok := r.URL.Query()[req.DBIssueID.GetMetaName()]; ok && len(vals) == 1 {
				var err error
				req.ID, err = strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			format := fmt.Sprintf("SELECT id, summary, type, COALESCE(fields, '{}'::JSONB) AS fields FROM %s.issues WHERE id = %d LIMIT 1", strings.ToUpper(req.Link), req.ID)
			rows, err := h.Queryx(format)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer rows.Close()

			issue := struct {
				types.DBIssue
				Type   int           `db:"type" json:"type"`
				Fields fields.Fields `db:"fields" json:"fields"`
			}{}
			for rows.Next() {
				rows.StructScan(&issue)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(issue)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
