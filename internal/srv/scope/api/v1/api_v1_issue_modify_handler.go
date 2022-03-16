package apiv1

import (
	"net/http"
	aur "prb/internal/srv/scope/common/auto_registerer"
)

type APIV1IssueModifyHandler struct {
	*aur.AutoRegistereesShared
}

func NewAPIV1IssueModifyHandler(sh *aur.AutoRegistereesShared) *APIV1IssueModifyHandler {
	return &APIV1IssueModifyHandler{
		AutoRegistereesShared: sh,
	}
}

func (h *APIV1IssueModifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}