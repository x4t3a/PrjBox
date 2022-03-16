package web

import (
	"html/template"
	"net/http"
	"path"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/types"

	"github.com/gorilla/mux"
)

type WebProjectManage struct {
	*aur.AutoRegistereesShared
}

func NewWebProjectManageHandler(sh *aur.AutoRegistereesShared) *WebProjectManage {
	return &WebProjectManage{
		AutoRegistereesShared: sh,
	}
}

func (h *WebProjectManage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("layout").Delims("[[", "]]").ParseFiles(
		path.Join(h.BasePath, "/web/prb/layout.html"),
		path.Join(h.BasePath, "/web/prb/util/project.manage.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	link, ok := vars["link"]
	if !ok {
		return
	}

	var name []string
	err = h.DB.Select(&name, "SELECT name FROM common.projects WHERE link=$1 LIMIT 1", link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		types.WebPageBaseHandlerData
		Name, Link string
	}{
		WebPageBaseHandlerData: types.WebPageBaseHandlerData{PageTitle: "manage link"},
		Name:                   name[0],
		Link:                   link,
	}

	tmpl.Execute(w, data)
}
