package web

import (
	"html/template"
	"net/http"
	"path"
	aur "prb/internal/srv/common/auto_registerer"
	"prb/internal/srv/common/types"
)

func Custom404Page(sh *aur.AutoRegistereesShared, w http.ResponseWriter, outLink, outName string) {
	w.WriteHeader(http.StatusNotFound)

	tmpl, err := template.New("layout").Delims("[[", "]]").ParseFiles(
		path.Join(sh.BasePath, "/web/prb/layout.html"),
		path.Join(sh.BasePath, "/web/prb/util/404.html"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := struct {
		types.WebPageBaseHandlerData
		OutLink, OutName string
	}{
		WebPageBaseHandlerData: types.WebPageBaseHandlerData{PageTitle: "404"},
		OutLink:                outLink,
		OutName:                outName,
	}

	_ = tmpl.Execute(w, data)
}
