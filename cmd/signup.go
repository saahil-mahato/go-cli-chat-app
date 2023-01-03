package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type User struct {
	Id       uint64
	Username string
	Password string
}

func init() {
	rootCmd.AddCommand(signupCmd)
}

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Signup",
	Long:  "Singup new user",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 || len(args[0]) < 1 || len(args[1]) < 1 {
			fmt.Println("Please enter username and password")
			return
		}

		newUser := &User{
			Username: args[0],
			Password: args[1],
		}
		_, err := db.Model(newUser).Insert()
		if err != nil {
			fmt.Println(err)
		}
	},
}
