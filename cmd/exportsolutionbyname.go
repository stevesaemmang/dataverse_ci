package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"stevesaemmang/dataverse_ci/dataverse"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var exportSolutionByNameCmd = &cobra.Command{
	Use:   "exportsolution",
	Short: "Exports solution by unique name.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tenantId, clientId, clientSecret, environment := checkConnectionFlags(cmd)
		solutionName, _ := cmd.Flags().GetString("solutionname")
		managed, _ := cmd.Flags().GetBool("managed")

		service := dataverse.InitService(tenantId, clientId, clientSecret, environment)

		fmt.Printf("Exporting solution...")

		byteArray := dataverse.ExportSolutionByName(service, solutionName, managed).ExportSolutionFile

		filename := solutionName

		if managed {
			filename = filename + "_managed"
		} else {
			filename = filename + "_unmanaged"
		}

		os.Mkdir("output", 0755)
		ioutil.WriteFile("output/"+filename+".zip", byteArray, 0644)
	},
}

func init() {
	rootCmd.AddCommand(exportSolutionByNameCmd)
	addConnectionFlags(exportSolutionByNameCmd)

	exportSolutionByNameCmd.Flags().StringP("solutionname", "s", "", "The unique name of the solution. (required)")
	exportSolutionByNameCmd.Flags().BoolP("managed", "m", false, "Specifies if the exported solution is managed.")
}
