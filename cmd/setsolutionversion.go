package cmd

import (
	"stevesaemmang/dataverse_ci/dataverse"

	"github.com/spf13/cobra"
)

var setSolutionVersionCmd = &cobra.Command{
	Use:   "setsolutionversion",
	Short: "Sets solution version.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tenantId, clientId, clientSecret, environment := checkConnectionFlags(cmd)
		solutionId, _ := cmd.Flags().GetString("solutionid")
		version, _ := cmd.Flags().GetString("version")

		service := dataverse.InitService(tenantId, clientId, clientSecret, environment)

		dataverse.SetSolutionVersion(service, solutionId, version)
	},
}

func init() {
	rootCmd.AddCommand(setSolutionVersionCmd)
	addConnectionFlags(setSolutionVersionCmd)

	setSolutionVersionCmd.Flags().StringP("solutionid", "s", "", "The id of the solution. (required)")
	setSolutionVersionCmd.Flags().StringP("version", "v", "", "The new version identifier. (required)")
}
