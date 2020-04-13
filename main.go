package main

import (
	"github.com/ysdy/Go_Learning/db"
	"github.com/ysdy/Go_Learning/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
