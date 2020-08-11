package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var rootCmd = &cobra.Command{
	Use:   "gkeyring",
	Short: "Access Gnome keyring from the cli",
	Long:  `Store and retrieve password from the Gnome keyring`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func main() {
	var setCmd = &cobra.Command{
		Use:   "set",
		Short: "Store a value in the keyring",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			app := args[0]
			user := args[1]
			password := args[2]
			// set password
			err := keyring.Set(app, user, password)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Retrieve a value in the keyring",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			app := args[0]
			user := args[1]
			// set password
			secret, err := keyring.Get(app, user)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(secret)
		},
	}
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(getCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
