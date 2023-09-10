package blockchain

import "database/sql"

type Blockchain struct {
	DB    *sql.DB
	index uint64
}
