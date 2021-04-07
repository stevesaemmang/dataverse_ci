package cmd

import (
	"fmt"
	"io/ioutil"
	"stevesaemmang/dataverse_ci/dataverse"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var importSolutionCmd = &cobra.Command{
	Use:   "importsolution",
	Short: "Imports solution.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tenantId, clientId, clientSecret, environment := checkConnectionFlags(cmd)
		solutionFile, _ := cmd.Flags().GetString("solutionfile")
		overwriteUnmanagedCustomizations, _ := cmd.Flags().GetBool("overwriteunmanaged")
		publishWorkflows, _ := cmd.Flags().GetBool("publishworkflows")

		service := dataverse.InitService(tenantId, clientId, clientSecret, environment)

		fmt.Println("Importing solution...")

		byteArray, err := ioutil.ReadFile(solutionFile)

		if err != nil {
			panic(err)
		}

		dataverse.ImportSolution(service, byteArray, overwriteUnmanagedCustomizations, publishWorkflows)
	},
}

func init() {
	rootCmd.AddCommand(importSolutionCmd)
	addConnectionFlags(importSolutionCmd)

	importSolutionCmd.Flags().StringP("solutionfile", "s", "", "The unique name of the solution. (required)")
	importSolutionCmd.Flags().Bool("overwriteunmanaged", false, "Specifies if the exported solution is managed.")
	importSolutionCmd.Flags().Bool("publishworkflows", false, "Specifies if the exported solution is managed.")
}
