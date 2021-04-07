package dataverse

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Solution struct {
	SolutionId   string `json:"solutionid"`
	UniqueName   string `json:"uniquename"`
	FriendlyName string `json:"friendlyname"`
	IsManaged    bool   `json:"ismanaged"`
	Version      string `json:"version"`
	Publisher    struct {
		FriendlyName string `json:"friendlyname"`
	} `json:"publisherid"`
}

func GetSolutionsByName(service IService, solutionName string) (ret struct {
	List []Solution `json:"value"`
}) {
	response := service.Get(fmt.Sprintf("solutions?$filter=uniquename%%20eq%%20%%27%s%%27&$expand=publisherid", solutionName))

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&ret)

	return
}

func GetSolutions(service IService) (ret struct {
	List []Solution `json:"value"`
}) {
	response := service.Get("solutions?$expand=publisherid")

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&ret)

	return
}

func GetSolutionsById(service IService, solutionId string) (ret Solution) {
	response := service.Get(fmt.Sprintf("solutions(%s)?$expand=publisherid", solutionId))

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&ret)

	return
}

func SetSolutionVersion(service IService, solutionId string, version string) {
	payload, _ := json.Marshal(struct {
		Version string `json:"version"`
	}{
		Version: version,
	})

	service.Patch(fmt.Sprintf("solutions(%s)", solutionId), payload)

	output, _ := json.Marshal(GetSolutionsById(service, solutionId))
	fmt.Println(string(output))

	return
}

func ExportSolutionByName(service IService, solutionName string, managed bool) (ret struct {
	ExportSolutionFile []byte `json:"ExportSolutionFile"`
}) {
	payload, _ := json.Marshal(struct {
		SolutionName string
		Managed      bool
	}{
		SolutionName: solutionName,
		Managed:      managed,
	})

	response := service.Post("ExportSolution", payload)

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&ret)

	return
}

func ImportSolution(service IService, solutionFile []byte, overwriteUnmanagedCustomizations bool, publishWorkflows bool) {
	payload, _ := json.Marshal(struct {
		OverwriteUnmanagedCustomizations bool
		PublishWorkflows                 bool
		CustomizationFile                []byte
		ImportJobId                      string
	}{
		OverwriteUnmanagedCustomizations: overwriteUnmanagedCustomizations,
		PublishWorkflows:                 publishWorkflows,
		CustomizationFile:                solutionFile,
		ImportJobId:                      fmt.Sprintf("%v", uuid.New()),
	})

	service.Post("ImportSolution", payload)

	return
}
