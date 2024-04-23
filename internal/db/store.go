package db

import "database/sql"

type SQLStore struct {
	*Queries
	ConnPool *sql.DB
}

func NewStore(connPool *sql.DB) Querier {
	return &SQLStore{
		Queries:  New(connPool),
		ConnPool: connPool,
	}
}
