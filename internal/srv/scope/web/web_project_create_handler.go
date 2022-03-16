package web

import (
	"html/template"
	"net/http"
	"path"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/types"
)

type WebProjectCreate struct{
	*aur.AutoRegistereesShared
}

func NewWebProjectCreateHandler(sh *aur.AutoRegistereesShared) *WebProjectCreate {
	return &WebProjectCreate{
		AutoRegistereesShared: sh,
	}
}

func (h *WebProjectCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("layout").Delims("[[", "]]").ParseFiles(
		path.Join(h.BasePath, "/web/prb/layout.html"),
		path.Join(h.BasePath, "/web/prb/util/project.create.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		types.WebPageBaseHandlerData
	}{
		WebPageBaseHandlerData: types.WebPageBaseHandlerData{PageTitle: "create project"},
	}

	tmpl.Execute(w, data)
}