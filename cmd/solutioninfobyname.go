package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"stevesaemmang/dataverse_ci/dataverse"

	"github.com/spf13/cobra"
)

var solutionByNameCmd = &cobra.Command{
	Use:   "solutionbyname",
	Short: "Gets solution info by unique name.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tenantId, clientId, clientSecret, environment := checkConnectionFlags(cmd)
		solutionName, _ := cmd.Flags().GetString("solutionname")

		service := dataverse.InitService(tenantId, clientId, clientSecret, environment)

		solutions := dataverse.GetSolutionsByName(service, solutionName)

		if len(solutions.List) == 0 {
			fmt.Printf("Solution not found.")
			os.Exit(1)
		}

		output, _ := json.Marshal(solutions.List[0])
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(solutionByNameCmd)
	addConnectionFlags(solutionByNameCmd)

	solutionByNameCmd.Flags().StringP("solutionname", "s", "", "The unique name of the solution. (required)")
}
