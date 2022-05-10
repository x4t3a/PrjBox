package db

import "database/sql"

type DBExecutor interface {
	Exec(query string, args ...any) (sql.Result, error)
}

type DBQuery struct {
	Query string
	Args  []any
}

type DBQueries []DBQuery

func ExecChain(dbe DBExecutor, queries DBQueries) (err error) {
	for _, query := range queries {
		_, err = dbe.Exec(query.Query, query.Args...)
		if err != nil {
			break
		}
	}

	return
}
