package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Long:  "Login user",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 || len(args[0]) < 1 || len(args[1]) < 1 {
			fmt.Println("Please enter username and password")
			return
		}

		user := new(User)
		err := db.Model(user).
			Where("username = ? AND password = ?", args[0], args[1]).
			Select()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Welcome to the chat room")
		}
	},
}
