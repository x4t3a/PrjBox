package apiv1

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"prb/internal/srv/scope/api/v1/fields"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/stub"
	"prb/internal/srv/scope/common/types"
)

type APIV1IssueCreateHandler struct {
	*aur.AutoRegistereesShared
}

func NewAPIV1IssueCreateHandler(sh *aur.AutoRegistereesShared) *APIV1IssueCreateHandler {
	return &APIV1IssueCreateHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *APIV1IssueCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type apiIssueCreateRequest struct {
		Project string `db:"project" json:"project"`
		types.DBIssue
	}

	if stub.UserHasRights() {
		switch r.Method {
		case "POST", "PUT":
			var req apiIssueCreateRequest
			if _, ok := r.URL.Query()["multipart"]; ok {
				err := r.ParseMultipartForm(1024)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				// mp.Decode(req, "json", r) // TODO
				req.Project = r.FormValue("project")
				req.Summary = r.FormValue("summary")
			} else {
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&req)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}

			tx, err := h.DB.BeginTxx(context.Background(), &sql.TxOptions{
				Isolation: 0,
				ReadOnly:  false,
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			reply := struct {
				Project string `json:"project"`
				ID      int64  `json:"id"`
				Summary string `json:"summary"`
			}{
				Project: req.Project,
			}

			defer func() {
				if err == nil {
					err = tx.Commit()
				} else {
					err = tx.Rollback()
				}

				if err == nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusAccepted)
					json.NewEncoder(w).Encode(reply)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			fields := fields.NewFieldsAuto()

			rows, err := tx.Query(fmt.Sprintf("INSERT INTO %s.issues(summary, fields) VALUES($1, $2) RETURNING id, summary", req.Project), req.Summary, fields)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			for rows.Next() {
				err = rows.Scan(&reply.ID, &reply.Summary)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
