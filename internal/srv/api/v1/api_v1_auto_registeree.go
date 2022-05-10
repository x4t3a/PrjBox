package apiv1

import (
	apiv1mini "prb/internal/srv/api/v1/mini"
	aur "prb/internal/srv/common/auto_registerer"
	"prb/internal/srv/common/routes"
)

type APIV1AutoRegisteree struct {
	pathPrefix string
}

func NewAPIV1AutoRegisteree(pathPrefix string) *APIV1AutoRegisteree {
	return &APIV1AutoRegisteree{pathPrefix: pathPrefix}
}

func (ar *APIV1AutoRegisteree) AutoRegister(sh *aur.AutoRegistereesShared) {
	subrouter := sh.Router.PathPrefix(ar.GetPathPrefix()).Subrouter()

	routes.Routes(subrouter, NewAPIV1IssueCreateHandler(sh), "/i/create")
	routes.Routes(subrouter, NewAPIV1IssueDeleteHandler(sh), "/i/delete")
	routes.Routes(subrouter, NewAPIV1IssueGetHandler(sh), "/i/get")
	routes.Routes(subrouter, NewAPIV1IssueModifyHandler(sh), "/i/modify")
	routes.Routes(subrouter, NewAPIV1ProjectArchiveHandler(sh), "/p/archive")
	routes.Routes(subrouter, NewAPIV1ProjectCreateHandler(sh), "/p/create")
	routes.Routes(subrouter, NewAPIV1ProjectDeleteHandler(sh), "/p/delete")

	routes.Routes(subrouter, apiv1mini.NewAPIV1MiniProjectsListHandler(sh), "/mini/p/list")
	routes.Routes(subrouter, apiv1mini.NewAPIV1MiniProjectsListHandler(sh), "/mini/status/flow")
}

func (ar *APIV1AutoRegisteree) GetPathPrefix() string {
	return ar.pathPrefix
}
