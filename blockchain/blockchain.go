package blockchain

import "database/sql"

type Blockchain struct {
	DB    *sql.DB
	index uint64
}

type Block struct {
	CurrHash []byte
	Prevhash []byte
	Nonce    uint64
	difficalty
}
