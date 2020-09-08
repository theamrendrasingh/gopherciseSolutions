package main

import (
	"path/filepath"

	"amru.in/cli/db"

	"amru.in/cli/cmd"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.OpenDB(dbPath)

	if err != nil {
		panic(err)
	}

	cmd.RootCmd.Execute()
}
