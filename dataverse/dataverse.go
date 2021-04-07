package dataverse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type service struct {
	TenantId     string
	ClientId     string
	ClientSecret string
	Environment  string
	BearerToken  string
}

type IService interface {
	Get(query string) *http.Response
	Post(query string, payload []byte) *http.Response
	Patch(query string, payload []byte)
}

func InitService(tenantId string, clientId string, clientSecret string, environment string) IService {
	service := service{
		tenantId, clientId, clientSecret, environment, "",
	}

	service.BearerToken = auth(service).AccessToken
	return service
}

func (srv service) Get(query string) *http.Response {
	request := srv.buildRequest("GET", query, nil)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		panic(fmt.Sprintf("HTTP Request failed: %s", response.Status))
	}

	return response
}

func (srv service) Post(query string, payload []byte) *http.Response {
	request := srv.buildRequest("POST", query, payload)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", fmt.Sprintf("%v", len(payload)))

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 && response.StatusCode != 204 {
		if response.StatusCode == http.StatusInternalServerError {
			defer response.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(response.Body)
			bodyString := string(bodyBytes)

			panic(bodyString)
		}

		panic(fmt.Sprintf("HTTP Request failed: %s", response.Status))
	}

	return response
}

func (srv service) Patch(query string, payload []byte) {
	request := srv.buildRequest("PATCH", query, payload)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", fmt.Sprintf("%v", len(payload)))
	request.Header.Set("Accept", "*/*")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 204 {
		panic(fmt.Sprintf("HTTP Request failed: %s", response.Status))
	}

	return
}

func (srv service) buildRequest(verb string, query string, payload []byte) *http.Request {
	request, _ := http.NewRequest(verb, fmt.Sprintf("%s/api/data/v9.2/%s", srv.Environment, query), bytes.NewBuffer(payload))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", srv.BearerToken))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("OData-MaxVersion", "4.0")
	request.Header.Set("OData-Version", "4.0")

	return request
}

func auth(srv service) (ret struct {
	AccessToken string `json:"access_token"`
}) {
	response, err := http.PostForm(fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", srv.TenantId), url.Values{
		"resource":      {srv.Environment},
		"client_id":     {srv.ClientId},
		"client_secret": {srv.ClientSecret},
		"grant_type":    {"client_credentials"}})

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&ret)

	return
}
