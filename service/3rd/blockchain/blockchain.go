package blockchain

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const (
	getAccountAddressMethod = "getaccountaddress"
	dumpPrivKey             = "dumpprivkey"
)

type Blockchain struct {
	c        *http.Client
	endpoint string
}

func New(c *http.Client, endpoint string) *Blockchain {
	return &Blockchain{
		c:        c,
		endpoint: endpoint,
	}
}

type accountAddressResp struct {
	Result struct {
		PrivateKey     string `json:"PrivateKey"`
		PaymentAddress string `json:"PaymentAddress"`
		ReadonlyKey    string `json:"ReadonlyKey"`
	} `json:"Result"`
	Error *string `json:"Error"`
	ID    int     `json:"Id"`
}

func (b *Blockchain) GetAccountWallet(params string) (paymentAddress, readonlyKey, privKey string, err error) {
	paymentAddress, readonlyKey, err = b.GetAccountAddress(params)
	if err != nil {
		err = errors.Wrap(err, "b.GetAccountAddress")
		return
	}
	privKey, err = b.DumpPrivKey(paymentAddress)
	if err != nil {
		err = errors.Wrap(err, "b.DumpPrivKey")
		return
	}
	return
}

func (b *Blockchain) GetAccountAddress(params string) (paymentAddress, readonlyKey string, err error) {
	body, err := b.post(b.buildRequestPayload(getAccountAddressMethod, params))
	if err != nil {
		return "", "", errors.Wrap(err, "b.post")
	}

	var resp accountAddressResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", "", errors.Wrap(err, "json.Unmarshal")
	}

	paymentAddress = resp.Result.PaymentAddress
	readonlyKey = resp.Result.ReadonlyKey
	return

}

func (b *Blockchain) DumpPrivKey(params string) (string, error) {
	body, err := b.post(b.buildRequestPayload(dumpPrivKey, params))
	if err != nil {
		return "", errors.Wrap(err, "b.post")
	}

	var resp accountAddressResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", errors.Wrap(err, "json.Unmarshal")
	}
	return resp.Result.PrivateKey, nil
}

func (b *Blockchain) buildRequestPayload(method, params string) map[string]interface{} {
	return map[string]interface{}{
		"jsonrpc": "1.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}
}

func (b *Blockchain) post(args map[string]interface{}) ([]byte, error) {
	data, err := json.Marshal(args)
	if err != nil {
		return nil, errors.Wrap(err, "json.Marshal")
	}

	req, err := http.NewRequest(http.MethodPost, b.endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "b.c.Do")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	return body, nil
}
