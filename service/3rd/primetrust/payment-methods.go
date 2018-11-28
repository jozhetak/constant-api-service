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

func CreateNewPaymentMethod(contact *models.PaymentMethod) (*models.PaymentMethod, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(contact)

	apiUrl := fmt.Sprintf("%s/payment-methods", _apiPrefix)
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

	response := models.PaymentMethod{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func GetPaymentMethods() (*models.PaymentMethodsResponse, error) {
	apiUrl := fmt.Sprintf("%s/payment-methods", _apiPrefix)
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

	response := models.PaymentMethodsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func GetPaymentMethod(paymentMethodId string) (*models.PaymentMethod, error) {
	apiUrl := fmt.Sprintf("%s/payment-methods/%s", _apiPrefix, paymentMethodId)
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

	response := models.PaymentMethod{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func DeletePaymentMethod(paymentMethodId string) (error) {
	apiUrl := fmt.Sprintf("%s/payment-methods/%s", _apiPrefix, paymentMethodId)
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
