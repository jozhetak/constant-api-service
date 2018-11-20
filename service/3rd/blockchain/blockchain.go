package blockchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const (
	dumpPrivKeyMethod       = "dumpprivkey"
	getAccountAddressMethod = "getaccountaddress"

	// wallet methods
	listAccountsMethod           = "listaccounts"
	getAccountMethod             = "getaccount"
	encryptDataMethod            = "encryptdata"
	getBalanceByPrivateKeyMethod = "getbalancebyprivatekey"

	// custom token
	getlistcustomtokenbalance = "getlistcustomtokenbalance"
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
	body, err := b.post(b.buildRequestPayload(dumpPrivKeyMethod, params))
	if err != nil {
		return "", errors.Wrap(err, "b.post")
	}

	var resp accountAddressResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", errors.Wrap(err, "json.Unmarshal")
	}
	return resp.Result.PrivateKey, nil
}

func (b *Blockchain) buildRequestPayload(method string, params interface{}) map[string]interface{} {
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

func (b *Blockchain) walletAPI(method string, params interface{}) (interface{}, error) {
	body, err := b.post(b.buildRequestPayload(method, params))
	if err != nil {
		return nil, errors.Wrap(err, "b.post")
	}

	fmt.Printf("string(body) = %+v\n", string(body))

	var v interface{}
	if err := json.Unmarshal(body, &v); err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}
	return v, nil
}

func (b *Blockchain) ListAccounts(params interface{}) (interface{}, error) {
	return b.walletAPI(listAccountsMethod, params)
}

func (b *Blockchain) GetAccount(params interface{}) (interface{}, error) {
	return b.walletAPI(getAccountMethod, params)
}

func (b *Blockchain) EncryptData(pubKey string, params interface{}) (string, error) {
	resp, err := b.walletAPI(encryptDataMethod, []interface{}{pubKey, params})
	if err != nil {
		return "", errors.Wrap(err, "b.walletAPI")
	}

	v, ok := resp.(map[string]interface{})
	if !ok {
		return "", errors.New("invalid response from blockchain core api")
	}

	encrypted, ok := v["Result"].(string)
	if !ok {
		return "", nil
	}
	return encrypted, nil
}

func (b *Blockchain) GetBalanceByPrivateKey(privKey string) (uint64, error) {
	resp, err := b.walletAPI(getBalanceByPrivateKeyMethod, []interface{}{privKey})
	if err != nil {
		return 0, err
	}
	data := resp.(map[string]interface{})
	return uint64(data["Result"].(float64)), nil
}

func (b *Blockchain) GetListCustomTokenBalance(paymentAddress string) (*ListCustomTokenBalance, error) {
	resp, err := b.walletAPI(getlistcustomtokenbalance, []interface{}{paymentAddress})
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"].(map[string]interface{})
	_ = data["Error"]
	result := ListCustomTokenBalance{}
	resultRespStr, err := json.Marshal(resultResp)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(resultRespStr), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
