package apiv1

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	aur "prb/internal/srv/common/auto_registerer"
	"prb/internal/srv/common/db"
	"prb/internal/srv/common/stub"
	"prb/internal/srv/common/types"
	"strings"
)

type APIV1ProjectCreateHandler struct {
	*aur.AutoRegistereesShared
}

func NewAPIV1ProjectCreateHandler(sh *aur.AutoRegistereesShared) *APIV1ProjectCreateHandler {
	return &APIV1ProjectCreateHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *APIV1ProjectCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type apiProjectCreateRequest struct {
		types.DBProject
	}

	if stub.UserHasRights() {
		switch r.Method {
		case "POST", "PUT":
			var req apiProjectCreateRequest
			if _, ok := r.URL.Query()["multipart"]; ok {
				err := r.ParseMultipartForm(1024)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				// mp.Decode(req, "json", r) // TODO
				req.Name = r.FormValue("name")
				req.Link = r.FormValue("link")
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
				}

				if err != nil {
					tx.Rollback()
				}
			}()

			req.Link = strings.ToUpper(req.Link)

			err = db.ExecChain(tx,
				db.DBQueries{
					db.DBQuery{Query: "INSERT INTO common.projects(name, link) VALUES($1, $2)", Args: []any{req.Name, req.Link}},
					db.DBQuery{Query: fmt.Sprintf("CREATE SCHEMA %s", req.Link)},
					db.DBQuery{Query: fmt.Sprintf("CREATE TABLE %s.issues(id SERIAL, type INTEGER NOT NULL DEFAULT 0, summary TEXT NOT NULL, fields JSONB)", req.Link)},
				})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusAccepted)
		case "GET":
			// TODO
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
