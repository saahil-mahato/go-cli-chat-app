package cmd

import (
	"github.com/go-pg/pg"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "chat-app",
		Short: "Chat app",
		Long:  `Chat app is an application`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var db *pg.DB = pg.Connect(&pg.Options{
	Addr:     "127.0.0.1:5432",
	User:     "postgres",
	Password: "postgres",
	Database: "chatdatabase",
})
