package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	aur "prb/internal/srv/common/auto_registerer"
	"prb/internal/srv/common/types"
)

type WebIssueCreate struct {
	*aur.AutoRegistereesShared
}

func NewWebIssueCreateHandler(sh *aur.AutoRegistereesShared) *WebIssueCreate {
	return &WebIssueCreate{
		AutoRegistereesShared: sh,
	}
}

func (h *WebIssueCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("layout").Delims("[[", "]]").ParseFiles(
		path.Join(h.BasePath, "/web/prb/layout.html"),
		path.Join(h.BasePath, "/web/prb/util/issue.create.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://%s/api/v1/mini/p/list", h.APIInterface))
	if err != nil {
		Custom404Page(h.AutoRegistereesShared, w, "/prb/p", "projects")
		return
	}

	data := struct {
		types.WebPageBaseHandlerData
		Projects []types.DBProject `json:"projects"`
	}{
		WebPageBaseHandlerData: types.WebPageBaseHandlerData{PageTitle: "create issue"},
		Projects:               make([]types.DBProject, 0),
	}

	json.NewDecoder(resp.Body).Decode(&data.Projects)

	err = h.DB.Select(&data.Projects, "SELECT link FROM common.projects ORDER BY link ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
