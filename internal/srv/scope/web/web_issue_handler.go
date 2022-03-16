package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"

	"prb/internal/srv/scope/api/v1/fields"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/types"

	"github.com/gorilla/mux"
)

type WebIssueHandler struct {
	*aur.AutoRegistereesShared
}

func NewWebIssueHandler(sh *aur.AutoRegistereesShared) *WebIssueHandler {
	return &WebIssueHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *WebIssueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		project, id string
		ok          bool
	)
	vars := mux.Vars(r)
	if project, ok = vars["project"]; !ok {
		return
	}

	if id, ok = vars["id"]; !ok {
		return
	}

	type dbIssue struct {
		ID      int64         `json:"id"`
		Summary string        `json:"summary"`
		Fields  fields.Fields `json:"fields"`
	}

	resp, err := http.Get(fmt.Sprintf("http://%s/api/v1/i/get?link=%s&id=%s", h.APIInterface, strings.ToUpper(project), id))
	if err != nil {
		Custom404Page(h.AutoRegistereesShared, w, "/prb/p", "projects")
		return
	}

	data := struct {
		types.WebPageBaseHandlerData
		Project string
		DBIssue dbIssue
	}{
		WebPageBaseHandlerData: types.WebPageBaseHandlerData{PageTitle: project},
		Project:                project,
	}

	if err = json.NewDecoder(resp.Body).Decode(&data.DBIssue); err != nil {
		Custom404Page(h.AutoRegistereesShared, w, "/prb/p", "projects")
		return
	}

	tmpl, err := template.New("layout").Delims("[[", "]]").ParseFiles(
		path.Join(h.BasePath, "/web/prb/layout.html"),
		path.Join(h.BasePath, "/web/prb/issue.html"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = tmpl.Execute(w, data)
}
