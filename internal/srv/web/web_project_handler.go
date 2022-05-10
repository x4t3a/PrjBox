package web

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"prb/internal/srv/common/types"
	"strings"

	aur "prb/internal/srv/common/auto_registerer"

	"github.com/gorilla/mux"
)

type WebProjectHandler struct {
	*aur.AutoRegistereesShared
}

func NewWebProjectHandler(sh *aur.AutoRegistereesShared) *WebProjectHandler {
	return &WebProjectHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *WebProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if project, ok := vars["project"]; ok {
		project = strings.ToUpper(project)
		var prj types.DBProject
		if err := h.DB.Get(&prj, "SELECT name, link FROM common.projects WHERE link=$1 LIMIT 1", project); err != nil {
			Custom404Page(h.AutoRegistereesShared, w, "/prb/p", "projects")
			return
		}

		var issues []types.DBIssue
		if err := h.DB.Select(&issues, fmt.Sprintf("SELECT id, summary FROM %s.issues ORDER BY id ASC", project)); err != nil {
			Custom404Page(h.AutoRegistereesShared, w, "/prb/p", "projects")
			return
		}

		tmpl, err := template.New("layout").Delims("[[", "]]").ParseFiles(
			path.Join(h.BasePath, "/web/prb/layout.html"),
			path.Join(h.BasePath, "/web/prb/project.html"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := struct {
			types.WebPageBaseHandlerData
			types.DBProject
			Issues []types.DBIssue
		}{
			WebPageBaseHandlerData: types.WebPageBaseHandlerData{PageTitle: prj.Name},
			DBProject:              prj,
			Issues:                 issues,
		}

		_ = tmpl.Execute(w, data)

		return
	}

	Custom404Page(h.AutoRegistereesShared, w, "/prb/p", "projects")
}
