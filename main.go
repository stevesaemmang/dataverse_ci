package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	exportSolutionCommand := flag.NewFlagSet("exportsolution", flag.ExitOnError)

	tenantId := exportSolutionCommand.String("tid", "", "The azure tenant of the service principal. (required)")
	clientId := exportSolutionCommand.String("cid", "", "The service principals client id. (required)")
	clientSecret := exportSolutionCommand.String("secret", "", "A matching service principal secret. (required)")
	organizationURI := exportSolutionCommand.String("org", "", "The environment uri. (required)")

	solutionName := exportSolutionCommand.String("solution", "", "The solution name to export. (required)")
	managed := exportSolutionCommand.Bool("managed", false, "Set flag if you want a managed export.")

	if len(os.Args) < 2 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "exportsolution":
		exportSolutionCommand.Parse(os.Args[2:])

		if *tenantId == "" || *clientId == "" || *clientSecret == "" || *organizationURI == "" || *solutionName == "" {
			exportSolutionCommand.PrintDefaults()
			os.Exit(1)
		}

		fmt.Printf("Exporting solution...")

		byteArray := ExportSolution(*solutionName, *managed, *organizationURI, Auth(
			*tenantId,
			*clientId,
			*clientSecret,
			*organizationURI,
		).AccessToken).ExportSolutionFile

		filename := *solutionName

		if *managed {
			filename = filename + "_managed"
		} else {
			filename = filename + "_unmanaged"
		}

		ioutil.WriteFile("output/"+filename+".zip", byteArray, 0644)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
