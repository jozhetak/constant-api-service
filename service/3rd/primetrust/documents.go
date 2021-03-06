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

func (p *Primetrust) UploadDocument(document models.Document) (*models.DocumentResponse, error) {
	apiUrl := fmt.Sprintf("%s/uploaded-documents", p.Endpoint)

	jsonBytes, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(jsonBytes)
	log.Println(buffer)
	req, err := http.NewRequest("POST", apiUrl, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "multipart/form-data")
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

	response := models.DocumentResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}
