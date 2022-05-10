package autoregisterer

import (
	c "prb/internal/srv/common/config"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type AutoRegisteree interface {
	AutoRegister(*AutoRegistereesShared)
	GetPathPrefix() string
}

type AutoRegistereesShared struct {
	c.PRBConfig
	*sqlx.DB
	*mux.Router
}

func AutoRegister(shared *AutoRegistereesShared, autoRegisterees ...AutoRegisteree) {
	for _, au := range autoRegisterees {
		au.AutoRegister(shared)
	}
}
