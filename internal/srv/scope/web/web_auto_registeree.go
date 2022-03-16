package web

import (
	"net/http"
	"path"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/routes"
)

type WebAutoRegisteree struct {
	pathPrefix string
}

func NewWebAutoRegisteree(pathPrefix string) *WebAutoRegisteree {
	return &WebAutoRegisteree{pathPrefix: pathPrefix}
}

func (ar *WebAutoRegisteree) AutoRegister(sh *aur.AutoRegistereesShared) {
	sh.Router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(sh.BasePath, "/web/favicon.ico"))
	})

	sh.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Custom404Page(sh, w, "/prb/p", "projects")
	})

	subrouter := sh.Router.PathPrefix(ar.GetPathPrefix()).Subrouter()

	routes.Routes(subrouter, NewWebIssueCreateHandler(sh), "/util/i/create")
	routes.Routes(subrouter, NewWebIssueHandler(sh), "/i/{project}-{id:[0-9]+}")
	routes.Routes(subrouter, NewWebProjectHandler(sh), "/p/{project}")
	routes.Routes(subrouter, NewWebProjectsHandler(sh), "/p", "/")
	routes.Routes(subrouter, NewWebProjectCreateHandler(sh), "/util/p/create")
	routes.Routes(subrouter, NewWebProjectManageHandler(sh), "/p/manage/{link}")
	routes.Routes(subrouter, NewWebSprintHandler(sh), "/s/{project}/sprint/{sprint:[0-9]+}/")
}

func (ar *WebAutoRegisteree) GetPathPrefix() string {
	return ar.pathPrefix
}
