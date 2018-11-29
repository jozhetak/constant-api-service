package primetrust

import (
	"errors"
	"fmt"
	"net/http"
)

func AuthorizeContribution(contributionId string) (error) {
	apiUrl := fmt.Sprintf("%s/contributions/%s/sandbox/authorize", _apiPrefix, contributionId)
	req, err := http.NewRequest("POST", apiUrl, nil)
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

func AuthorizeDisbursement(disbursementId string) (error) {
	apiUrl := fmt.Sprintf("%s/disbursements/%s/sandbox/authorize", _apiPrefix, disbursementId)
	req, err := http.NewRequest("POST", apiUrl, nil)
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
