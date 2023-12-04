package main

import (
	"serverpackage/internal/auth"
	"serverpackage/internal/server"
)

func main() {
	auth.NewAuth()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
