package primetrust

import (
	"bytes"
	"io/ioutil"
	"github.com/ninjadotorg/constant-api-service/service/3rd/primetrust/models"
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
)

func CreateNewContribution(contact *models.Contribution) (*models.Contribution, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(contact)

	apiUrl := fmt.Sprintf("%s/contributions", _apiPrefix)
	req, err := http.NewRequest("POST", apiUrl, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", _authHeader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("%s: %s", res.Status, string(body)))
	}

	response := models.Contribution{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func GetContribution(contributionId string) (*models.Contribution, error) {
	apiUrl := fmt.Sprintf("%s/contributions/%s", _apiPrefix, contributionId)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", _authHeader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)

	response := models.Contribution{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func GetContributions() (*models.ContributionsResponse, error) {
	apiUrl := fmt.Sprintf("%s/contributions", _apiPrefix)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", _authHeader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)

	response := models.ContributionsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func DeleteContribution(contributionId string) (error) {
	apiUrl := fmt.Sprintf("%s/contributions/%s", _apiPrefix, contributionId)
	req, err := http.NewRequest("DELETE", apiUrl, nil)
	req.Header.Add("Authorization", _authHeader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	return nil
}
