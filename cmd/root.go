package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dataverse_ci",
	Short: "CLI tool for maintainance and CI/CD task in MS Dataverse.",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

func addConnectionFlags(command *cobra.Command) {
	command.Flags().String("tenantId", "", "The azure tenant the app id is registrated in. (required) (Format: 00000000-0000-0000-0000-000000000000)")
	command.Flags().String("clientId", "", "The app/client id. (required) (Format: 00000000-0000-0000-0000-000000000000)")
	command.Flags().String("clientSecret", "", "The matching secret of the specified client id. (required)")
	command.Flags().String("environment", "", "The url of the dataverse environment. (Example: https://myapp.crm4.dynamics.com)")
}

func checkConnectionFlags(command *cobra.Command) (string, string, string, string) {
	tenantId, _ := command.Flags().GetString("tenantId")
	clientId, _ := command.Flags().GetString("clientId")
	clientSecret, _ := command.Flags().GetString("clientSecret")
	environment, _ := command.Flags().GetString("environment")

	if tenantId == "" || clientId == "" || clientSecret == "" || environment == "" {
		fmt.Println(fmt.Errorf("Missing required connection arguments."))
		command.Help()
		os.Exit(0)
	}

	return tenantId, clientId, clientSecret, environment
}
