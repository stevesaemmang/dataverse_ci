package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ExportSolution(solutionName string, managed bool, organizationURI string, authToken string) (ret struct {
	ExportSolutionFile []byte `json:"ExportSolutionFile"`
}) {
	payload, _ := json.Marshal(struct {
		SolutionName string
		Managed      bool
	}{
		SolutionName: solutionName,
		Managed:      managed,
	})

	request, _ := http.NewRequest("POST", organizationURI+"/api/data/v9.2/ExportSolution", bytes.NewBuffer(payload))
	request.Header.Set("Authorization", "Bearer "+authToken)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&ret)

	return
}
