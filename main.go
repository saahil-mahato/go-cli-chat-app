package main

import (
	"chat-app/cmd"

	"github.com/go-pg/pg/v10"
)

var DB *pg.DB

func main() {
	cmd.Execute()
}
