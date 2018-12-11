package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ninjadotorg/constant-api-service/service/3rd/primetrust/models"
)

func (p *Primetrust) CreateNewWebhook(webhook *models.Webhook) (*models.Webhook, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(webhook)

	apiUrl := fmt.Sprintf("%s/webhook-configs", p.Endpoint)
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

	response := models.Webhook{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) UpdateWebhook(webhook *models.Webhook) (*models.Webhook, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(webhook)

	apiUrl := fmt.Sprintf("%s/webhook-configs/%s", p.Endpoint, webhook.Data.ID)
	log.Println(apiUrl)
	log.Println(jsonData)
	req, err := http.NewRequest("PATCH", apiUrl, jsonData)
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

	response := models.Webhook{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) GetWebhooks() (*models.WebhooksResponse, error) {
	apiUrl := fmt.Sprintf("%s/webhooks", p.Endpoint)
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

	response := models.WebhooksResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func (p *Primetrust) GetWebhook(webhookId string) (*models.Webhook, error) {
	apiUrl := fmt.Sprintf("%s/webhooks/%s", p.Endpoint, webhookId)
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

	response := models.Webhook{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}
