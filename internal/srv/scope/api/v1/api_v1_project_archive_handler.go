package apiv1

import (
	"net/http"
	aur "prb/internal/srv/scope/common/auto_registerer"
	"prb/internal/srv/scope/common/stub"
)


type APIV1ProjectArchiveHandler struct{
	*aur.AutoRegistereesShared
}

func NewAPIV1ProjectArchiveHandler(sh *aur.AutoRegistereesShared) *APIV1ProjectArchiveHandler {
	return &APIV1ProjectArchiveHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *APIV1ProjectArchiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stub.NotImplemented(w, r)
}