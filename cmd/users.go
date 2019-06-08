package cmd

import (
	"fmt"
	"os"

	"github.com/identixone/identixone-go/api/client"
	"github.com/identixone/identixone-go/utils"
	"github.com/spf13/cobra"
)

const (
	me             = "me"
	changePassword = "change-password"
	listTokens     = "list-tokens"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "users API",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("only one argument supported at time")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case me:
			c, err := client.NewClient()
			ifErrorExit(err)
			meOut, err := c.Users().Me()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			ifErrorExit(utils.PrettyPrint(meOut))
		case listTokens:
			c, err := client.NewClient()
			ifErrorExit(err)
			tokens, err := c.Users().ListTokens(nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			ifErrorExit(utils.PrettyPrint(tokens))
		}
	},
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
