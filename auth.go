package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func Auth(tenantId string, clientId string, clientSecret string, organizationURI string) (ret struct {
	AccessToken string `json:"access_token"`
}) {
	response, err := http.PostForm("https://login.microsoftonline.com/"+tenantId+"/oauth2/token", url.Values{
		"resource":      {organizationURI},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"grant_type":    {"client_credentials"}})

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&ret)

	return
}
