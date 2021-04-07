package cmd

import (
	"encoding/json"
	"fmt"
	"stevesaemmang/dataverse_ci/dataverse"

	"github.com/spf13/cobra"
)

var solutionListCmd = &cobra.Command{
	Use:   "solutionlist",
	Short: "Gets all installed solutions.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tenantId, clientId, clientSecret, environment := checkConnectionFlags(cmd)

		service := dataverse.InitService(tenantId, clientId, clientSecret, environment)

		solutions := dataverse.GetSolutions(service)

		output, _ := json.Marshal(solutions.List)
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(solutionListCmd)
	addConnectionFlags(solutionListCmd)
}
