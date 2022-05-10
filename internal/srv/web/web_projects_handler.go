package web

import (
	"html/template"
	"net/http"
	"path"
	"prb/internal/srv/common/types"

	aur "prb/internal/srv/common/auto_registerer"
)

type WebProjectsHandler struct {
	*aur.AutoRegistereesShared
}

func NewWebProjectsHandler(sh *aur.AutoRegistereesShared) *WebProjectsHandler {
	return &WebProjectsHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *WebProjectsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	data := struct {
		types.WebPageBaseHandlerData
		Projects []types.DBProject
	}{
		WebPageBaseHandlerData: types.WebPageBaseHandlerData{PageTitle: "projects"},
		Projects:               make([]types.DBProject, 0),
	}

	if err := h.Select(&data.Projects, "SELECT name, link FROM common.projects ORDER BY link ASC"); err != nil {
		Custom404Page(h.AutoRegistereesShared, w, "/prb/p", "projects")
		return
	}

	tmpl, err := template.New("layout").Delims("[[", "]]").ParseFiles(
		path.Join(h.BasePath, "/web/prb/layout.html"),
		path.Join(h.BasePath, "/web/prb/projects.html"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = tmpl.Execute(w, data)
	// _ = tmpl.ExecuteTemplate(w, "layout", data)
}
