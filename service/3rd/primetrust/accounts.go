package primetrust

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ninjadotorg/constant-api-service/service/3rd/primetrust/models"
)

func (p *Primetrust) GetAllAccounts() (*models.AccountsResponse, error) {
	apiUrl := fmt.Sprintf("%s/accounts", p.Endpoint)
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

	response := models.AccountsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) GetAccountByID(accountId string) (*models.Account, error) {
	apiUrl := fmt.Sprintf("%s/accounts/%s", p.Endpoint, accountId)
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

	response := models.Account{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}
