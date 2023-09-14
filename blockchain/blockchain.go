package blockchain

import (
	"crypto/rsa"
	"database/sql"
	"os"
	"time"
)

type Blockchain struct {
	DB    *sql.DB
	index uint64
}

type Block struct {
	CurrHash     []byte
	Prevhash     []byte
	Nonce        uint64
	Difficulty   uint8
	Miner        string
	Signature    []byte
	TimeStamp    string
	Transactions []Transaction
	Mapping      map[string]uint64
}

type Transaction struct {
	RandBytes []byte
	PrevBlock []byte
	Sender    string
	Receiver  string
	Value     uint64
	ToStorage uint64
	CurrHash  []byte
	Signature []byte
}

type User struct {
	PrivateKey *rsa.PrivateKey
}

const (
		CREATE_TABLE = `
CREATE TABLE BLOCKCHAIN (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Hash VARCHAR(44) UNIQUE,
    Block TEXT
)
`
)

const (
	GENESIS_BLOCK = "GENESIS-BLOCK"
	STORAGE_VALUE = 100
	GENESIS_REWARD = 100
	STORAGE_CHAIN = "STORAGE-CHAIN"
)

func NewChain(filename, receiver string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	file.Close()

	db, err := sql.Open('sqlite3', filename)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err := db.Exec(CREATE_TABLE)
	chain := &Blockchain{
		DB: db,
	}
	genesis := &Block{
		CurrHash: []byte,
		Mapping: make(map[string]uint64),
		Miner: receiver,
		TimeStamp: time.Now().Format(time.RFC3339),
	}
	genesis.Mapping[STORAGE_CHAIN] = STORAGE_VALUE
	genesis.Mapping[receiver] = GENESIS_REWARD
	chain.AddBlock(genesis)
	return nil
}

func (chain *Blockchain) AddBlock(block *Block){
	chain.index += 1
	chain.DB.Exec("INSERT INTO Blockchain (Hash, Block) VALUES ($1, $2)", 
		Base64Encode(block.CurrHash), 
		SerializeBlock(block),
	)
}

func SerializeBlock(block *Block) any {
	
}

func Base64Encode(hash []byte) any {
	
}

