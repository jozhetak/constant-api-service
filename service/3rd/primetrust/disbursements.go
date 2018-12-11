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

func (p *Primetrust) CreateDisbursement(disbursement *models.Disbursement) (*models.Disbursement, error) {
	jsonData := new(bytes.Buffer)
	err := json.NewEncoder(jsonData).Encode(disbursement)
	if err != nil {
		return nil, err
	}

	apiUrl := fmt.Sprintf("%s/disbursements", p.Endpoint)
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

	response := models.Disbursement{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) GetDisbursementByID(disbursementId string) (*models.Disbursement, error) {
	apiUrl := fmt.Sprintf("%s/disbursements/%s", p.Endpoint, disbursementId)
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

	response := models.Disbursement{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) GetAllDisbursements() (*models.DisbursementsResponse, error) {
	apiUrl := fmt.Sprintf("%s/disbursements", p.Endpoint)
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

	response := models.DisbursementsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) DeleteDisbursement(disbursementId string) error {
	apiUrl := fmt.Sprintf("%s/disbursements/%s", p.Endpoint, disbursementId)
	req, err := http.NewRequest("DELETE", apiUrl, nil)
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

func (p *Primetrust) AuthorizeDisbursement(disbursementId string) error {
	apiUrl := fmt.Sprintf("%s/disbursements/%s/sandbox/authorize", p.Endpoint, disbursementId)
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
