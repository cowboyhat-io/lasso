package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// show.goCmd represents the show.go command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "shows the environment variables that have been loaded from the lasso.env file",
	Long:  `Useful for checking credentials.`,
	Run: func(cmd *cobra.Command, args []string) {
		show()
	},
}

func show() {
	fmt.Printf("Project:%s\n", viper.GetString("PROJECT"))
	fmt.Printf("Commenter:%s\n", viper.GetString("COMMENTER"))
	fmt.Printf("OrgURL:%s\n", viper.GetString("ORG_URL"))
	fmt.Printf("PersonalAccessToken:%s\n", viper.GetString("PERSONAL_ACCESS_TOKEN"))
	fmt.Printf("Comment:%s\n", viper.GetString("COMMENT"))

}

func init() {
	rootCmd.AddCommand(showCmd)
}
