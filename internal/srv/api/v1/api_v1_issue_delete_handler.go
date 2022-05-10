package apiv1

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	aur "prb/internal/srv/common/auto_registerer"
	"prb/internal/srv/common/stub"
	"strconv"
)

type APIV1IssueDeleteHandler struct {
	*aur.AutoRegistereesShared
}

func NewAPIV1IssueDeleteHandler(sh *aur.AutoRegistereesShared) *APIV1IssueDeleteHandler {
	return &APIV1IssueDeleteHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *APIV1IssueDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type apiIssueDeleteRequest struct {
		Project string `db:"project" json:"project"`
		ID      int64  `db:"id"      json:"id"`
	}

	if stub.UserHasRights() {
		switch r.Method {
		case "POST", "PUT":
			var req apiIssueDeleteRequest
			if _, ok := r.URL.Query()["multipart"]; ok {
				err := r.ParseMultipartForm(1024)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				// mp.Decode(req, "json", r) // TODO
				req.Project = r.FormValue("project")
				if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
					req.ID = int64(id)
				}
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

			defer func() {
				if err == nil {
					err = tx.Commit()
				} else {
					err = tx.Rollback()
				}

				if err == nil {
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()
			_, err = tx.Exec(fmt.Sprintf("DELETE FROM %s.issues WHERE id = $1", req.Project), req.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
