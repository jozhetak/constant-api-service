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

func CreateNewUser(user *models.User) (*models.User, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(user)

	apiUrl := fmt.Sprintf("%s/users", _apiPrefix)
	req, err := http.NewRequest("POST", apiUrl, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", _authHeader)

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

	response := models.User{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func GetUser(userId string) (*models.User, error) {
	apiUrl := fmt.Sprintf("%s/users/%s", _apiPrefix, userId)
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := models.User{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func GetUsers() (*models.UsersResponse, error) {
	apiUrl := fmt.Sprintf("%s/users", _apiPrefix)
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := models.UsersResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func GetCurrentUser() (*models.User, error) {
	apiUrl := fmt.Sprintf("%s/users/current", _apiPrefix)
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := models.User{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func UpdateUserEmail(userId string, email string) (*models.User, error) {
	user := models.User{
		Data: models.UserData{
			Type: models.UserType,
			Attributes: models.UserAttributes{
				Email: email,
			},
		},
	}

	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(user)

	apiUrl := fmt.Sprintf("%s/users/%s", _apiPrefix, userId)
	req, err := http.NewRequest("PATCH", apiUrl, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", _authHeader)

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

	response := models.User{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func UpdateUserPassword(userId string, currentPassword string, password string) (*models.User, error) {
	user := models.User{
		Data: models.UserData{
			Type: models.UserType,
			Attributes: models.UserAttributes{
				CurrentPassword: currentPassword,
				Password:        password,
			},
		},
	}

	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(user)

	apiUrl := fmt.Sprintf("%s/users/%s/password", _apiPrefix, userId)
	req, err := http.NewRequest("PATCH", apiUrl, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", _authHeader)

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

	response := models.User{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}
