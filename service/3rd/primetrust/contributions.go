package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ninjadotorg/constant-api-service/service/3rd/primetrust/models"
)

func (p *Primetrust) CreateContribution(contribution *models.Contribution) (*models.Contribution, error) {
	jsonData := new(bytes.Buffer)
	err := json.NewEncoder(jsonData).Encode(contribution)
	if err != nil {
		return nil, err
	}

	apiUrl := fmt.Sprintf("%s/contributions", p.Endpoint)
	req, err := http.NewRequest("POST", apiUrl, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", p.Authorization)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("%s: %s", res.Status, string(body)))
	}

	response := models.Contribution{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) GetAllContributions() (*models.ContributionsResponse, error) {
	apiUrl := fmt.Sprintf("%s/contributions", p.Endpoint)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", p.Authorization)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := models.ContributionsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) AuthorizeContribution(contributionId string) error {
	apiUrl := fmt.Sprintf("%s/contributions/%s/sandbox/authorize", p.Endpoint, contributionId)
	req, err := http.NewRequest("POST", apiUrl, nil)
	req.Header.Add("Authorization", p.Authorization)

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
