package apiv1

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/db"
	"prb/internal/srv/scope/common/stub"
	"prb/internal/srv/scope/common/types"
	"strings"
)

type APIV1ProjectDeleteHandler struct {
	*aur.AutoRegistereesShared
}

func NewAPIV1ProjectDeleteHandler(sh *aur.AutoRegistereesShared) *APIV1ProjectDeleteHandler {
	return &APIV1ProjectDeleteHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *APIV1ProjectDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := struct {
		types.DBProjectLink
	}{}

	if stub.UserHasRights() {
		switch r.Method {
		case "POST", "PUT":
			if _, ok := r.URL.Query()["multipart"]; ok {
				err := r.ParseMultipartForm(1024)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				req.Link = r.FormValue("link")
			} else {
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&req)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
		case "DELETE":
			if vals, ok := r.URL.Query()["link"]; ok && len(vals) > 0 {
				req.Link = vals[0]
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
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
			db.DBQuery{Query: "DELETE FROM common.projects WHERE link=$1", Args: []any{req.Link}},
			db.DBQuery{Query: fmt.Sprintf("DROP SCHEMA %s CASCADE", req.Link)},
		})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
