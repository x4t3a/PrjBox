package apiv1

import (
	apiv1mini "prb/internal/srv/scope/api/v1/mini"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/routes"
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

	routes.Routes(subrouter, apiv1mini.NewAPIV1IssueCreateHandler(sh), "/mini/p/list")
}

func (ar *APIV1AutoRegisteree) GetPathPrefix() string {
	return ar.pathPrefix
}
