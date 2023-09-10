package main

import (
	nt "./network"
)

const (
	TO_UPPER = iota + 1
	TO_LOWER
)

const (
	ADDRESS = ":8080"
)

func main() {
	go nt.Listen(ADDRESS, handleServer)
}

func handleServer(conn nt.Conn, pack *nt.Package) {
	nt.Handle(TO_UPPER, conn, pack, handleToUpper)
}

func handleToUpper(pack *nt.Package)
