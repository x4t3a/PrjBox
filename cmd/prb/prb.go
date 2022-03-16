package main

import (
	"log"
	"net/http"
	apiv1 "prb/internal/srv/scope/api/v1"
	aur "prb/internal/srv/scope/common/auto_registerer"
	c "prb/internal/srv/scope/common/config"
	"prb/internal/srv/scope/web"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connectDB() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", "host=192.168.1.83 user=ad dbname=prb sslmode=disable")
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	shared := &aur.AutoRegistereesShared{
		PRBConfig: c.NewPRBConfig(),
		Router:    mux.NewRouter(),
		DB:        db,
	}

	var listenInterface string

	if shared.PRBConfig.IsAPI() {
		apiv1.NewAPIV1AutoRegisteree("/api/v1/").AutoRegister(shared)
		listenInterface = shared.APIInterface
	}

	if shared.PRBConfig.IsWeb() {
		web.NewWebAutoRegisteree("/prb/").AutoRegister(shared)
		listenInterface = shared.WebInterface
	}

	serveErr := http.ListenAndServe(listenInterface, shared.Router)
	log.Fatal(serveErr)
}
