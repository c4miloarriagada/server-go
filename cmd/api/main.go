package main

import (
	"serverpackage/internal/auth"
	db "serverpackage/internal/database"
	"serverpackage/internal/server"
)

func main() {

	auth.NewAuth()

	db.ConnectDatabase()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
