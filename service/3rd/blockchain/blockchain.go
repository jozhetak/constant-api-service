package blockchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/serializers"
)

const (
	DumpPrivKeyMethod = "dumpprivkey"
	GetAccountAddress = "getaccountaddress"
	LoanParams        = "loanparams"

	// wallet methods
	ListAccountsMethod           = "listaccounts"
	GetAccount                   = "getaccount"
	EncryptData                  = "encryptdata"
	GetBalanceByPrivateKeyMethod = "getbalancebyprivatekey"
	GetBalanceByPaymentAddress   = "getbalancebypaymentaddress"

	// tx
	CreateAndSendTransaction            = "createandsendtransaction"
	CreateAndSendCustomTokenTransaction = "createandsendcustomtokentransaction"
	GetTransactionByHash                = "gettransactionbyhash"
	CreateAndSendLoanRequest            = "createandsendloanrequest"
	createandsendloanresponse           = "createandsendloanresponse"
	CreateAndSendLoanPayment            = "createandsendloanpayment"
	CreateAndSendLoanWithdraw           = "createandsendloanwithdraw"

	// custom token
	GetListCustomTokenBalance = "getlistcustomtokenbalance"

	// voting
	GetBondTypes                         = "getbondtypes"
	GetGOVParams                         = "getgovparams"
	GetDCBParams                         = "getdcbparams"
	CreateAndSendVoteDCBBoardTransaction = "createandsendvotedcbboardtransaction"
	CreateAndSendVoteGOVBoardTransaction = "createandsendvotegovboardtransaction"
)

var (
	errTxHashNotExists = errors.New("tx hash does not exist")
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
	body, err := b.post(b.buildRequestPayload(GetAccountAddress, params))
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
	body, err := b.post(b.buildRequestPayload(DumpPrivKeyMethod, params))
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
		return nil, errors.Wrapf(err, "b.c.Do: %q", req.URL.String())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	return body, nil
}

func (b *Blockchain) blockchainAPI(method string, params interface{}) (interface{}, error) {
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
	return b.blockchainAPI(ListAccountsMethod, params)
}

func (b *Blockchain) GetAccount(params interface{}) (interface{}, error) {
	return b.blockchainAPI(GetAccount, params)
}

func (b *Blockchain) EncryptData(pubKey string, params interface{}) (string, error) {
	resp, err := b.blockchainAPI(EncryptData, []interface{}{pubKey, params})
	if err != nil {
		return "", errors.Wrap(err, "b.blockchainAPI")
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
	resp, err := b.blockchainAPI(GetBalanceByPrivateKeyMethod, []interface{}{privKey})
	if err != nil {
		return 0, err
	}
	data := resp.(map[string]interface{})
	return uint64(data["Result"].(float64)), nil
}

func (b *Blockchain) GetBalanceByPaymentAddress(paymentAddress string) (uint64, error) {
	resp, err := b.blockchainAPI(GetBalanceByPaymentAddress, []interface{}{paymentAddress})
	if err != nil {
		return 0, err
	}
	data := resp.(map[string]interface{})
	return uint64(data["Result"].(float64)), nil
}

func (b *Blockchain) GetListCustomTokenBalance(paymentAddress string) (*ListCustomTokenBalance, error) {
	resp, err := b.blockchainAPI(GetListCustomTokenBalance, []interface{}{paymentAddress})
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

func (b *Blockchain) CreateAndSendConstantTransaction(prvKey string, req serializers.WalletSend) (string, error) {
	param := []interface{}{prvKey, req.PaymentAddresses, -1, 8}
	resp, err := b.blockchainAPI(CreateAndSendTransaction, param)
	if err != nil {
		return "", errors.Wrap(err, "b.blockchainAPI")
	}
	data := resp.(map[string]interface{})
	txID, ok := data["Result"].(string)
	if !ok {
		return "", errors.Errorf("couldn't get txID: param: %+v, resp: %+v", param, data)
	}
	return txID, nil
}

func (b *Blockchain) SendCustomTokenTransaction(prvKey string, req serializers.WalletSend) error {
	param := []interface{}{prvKey, -1, 8}
	tokenData := map[string]interface{}{}
	tokenData["TokenID"] = req.TokenID
	tokenData["TokenTxType"] = 1
	tokenData["TokenName"] = req.TokenName
	tokenData["TokenSymbol"] = req.TokenSymbol
	tokenData["TokenReceivers"] = req.PaymentAddresses
	param = append(param, tokenData)
	resp, err := b.blockchainAPI(CreateAndSendCustomTokenTransaction, param)
	if err != nil {
		return errors.Wrap(err, "b.blockchainAPI")
	}

	var (
		data       = resp.(map[string]interface{})
		resultResp = data["Result"]
		errorResp  = data["Error"]
	)
	if errorResp != nil {
		return errors.Errorf("couldn't get result from response data: %+v", errorResp)
	}
	if resultResp == nil {
		return errors.Errorf("couldn't get result from response: req: %+v, data: %+v", param, data)
	}
	return nil
}

func (b *Blockchain) GetTxByHash(txHash string) (*TransactionDetail, error) {
	param := []interface{}{txHash}
	resp, err := b.blockchainAPI(GetTransactionByHash, param)
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"].(map[string]interface{})
	if resultResp["Hash"] == nil {
		return nil, errTxHashNotExists
	}
	result := TransactionDetail{}
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

func (b *Blockchain) CreateAndSendLoanRequest(prvKey string, request serializers.LoanRequest) (*string, error) {
	param := []interface{}{prvKey, 0, request}
	resp, err := b.blockchainAPI(CreateAndSendLoanRequest, param)
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"]
	if resultResp == nil {
		return nil, errors.New("Fail")
	}
	txID := resultResp.(string)
	return &txID, nil
}

func (b *Blockchain) CreateAndSendLoanWithdraw(prvKey string, request serializers.LoanWithdraw) (*string, error) {
	param := []interface{}{prvKey, 0, request}
	resp, err := b.blockchainAPI(CreateAndSendLoanWithdraw, param)
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"]
	if resultResp == nil {
		return nil, errors.New("Fail")
	}
	txID := resultResp.(string)
	return &txID, nil
}

func (b *Blockchain) CreateAndSendLoanPayment(prvKey string, request serializers.LoanPayment) (*string, error) {
	param := []interface{}{prvKey, 0, request}
	resp, err := b.blockchainAPI(CreateAndSendLoanPayment, param)
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"]
	if resultResp == nil {
		return nil, errors.New("Fail")
	}
	txID := resultResp.(string)
	return &txID, nil
}

func (b *Blockchain) GetLoanParams() ([]interface{}, error) {
	param := []interface{}{}
	resp, err := b.blockchainAPI(LoanParams, param)
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"]
	if resultResp == nil {
		return nil, errors.New("Fail")
	}
	return resultResp.([]interface{}), nil
}

func (b *Blockchain) GetBondTypes() ([]interface{}, error) {
	param := []interface{}{}
	resp, err := b.blockchainAPI(GetBondTypes, param)
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"]
	if resultResp == nil {
		return nil, errors.New("Fail")
	}
	return resultResp.([]interface{}), nil
}

func (b *Blockchain) GetGOVParams() ([]interface{}, error) {
	param := []interface{}{}
	resp, err := b.blockchainAPI(GetGOVParams, param)
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"]
	if resultResp == nil {
		return nil, errors.New("Fail")
	}
	return resultResp.([]interface{}), nil
}

func (b *Blockchain) GetDCBParams() ([]interface{}, error) {
	param := []interface{}{}
	resp, err := b.blockchainAPI(GetDCBParams, param)
	if err != nil {
		return nil, err
	}
	data := resp.(map[string]interface{})
	resultResp := data["Result"]
	if resultResp == nil {
		return nil, errors.New("Fail")
	}
	return resultResp.([]interface{}), nil
}

func (b *Blockchain) CreateAndSendVoteGOVBoardTransaction(privkey string, voteAmount uint64) error {
	param := []interface{}{privkey, -1, 8}
	tokenData := map[string]interface{}{}
	tokenData["TokenID"] = [32]byte{5} // DCB voting token
	tokenData["TokenTxType"] = 1
	tokenData["TokenName"] = ""   // TODO 0xjackalope get from stability
	tokenData["TokenSymbol"] = "" // TODO 0xjackalope get from stability
	govAccount := make(map[string]uint64)
	govAccount["1Uv1fjA1FjsLTp37i1j5ZVpghx3maaX6YM5WQkbtrJr26FyGwxKznAM7ZRN2AsE4iHwNjiWGLbcUt2JudBBek18cB5YV22EJ38PjcXqza"] = voteAmount
	tokenData["TokenReceivers"] = govAccount
	param = append(param, tokenData)
	param = append(param)
	resp, err := b.blockchainAPI(CreateAndSendVoteGOVBoardTransaction, param)
	if err != nil {
		return errors.Wrap(err, "b.blockchainAPI")
	}

	var (
		data       = resp.(map[string]interface{})
		resultResp = data["Result"]
		errorResp  = data["Error"]
	)
	if errorResp != nil {
		return errors.Errorf("couldn't get result from response data: %+v", errorResp)
	}
	if resultResp == nil {
		return errors.Errorf("couldn't get result from response: req: %+v, data: %+v", param, data)
	}
	return nil
}

func (b *Blockchain) CreateAndSendVoteDCBBoardTransaction(privkey string, voteAmount uint64) error {
	param := []interface{}{privkey, -1, 8}
	tokenData := map[string]interface{}{}
	tokenData["TokenID"] = [32]byte{6} // DCB voting token
	tokenData["TokenTxType"] = 1
	tokenData["TokenName"] = ""   // TODO 0xjackalope get from stability
	tokenData["TokenSymbol"] = "" // TODO 0xjackalope get from stability
	govAccount := make(map[string]uint64)
	govAccount["1Uv1fjA1FjsLTp37i1j5ZVpghx3maaX6YM5WQkbtrJr26FyGwxKznAM7ZRN2AsE4iHwNjiWGLbcUt2JudBBek18cB5YV22EJ38PjcXqza"] = voteAmount
	tokenData["TokenReceivers"] = govAccount
	param = append(param, tokenData)
	param = append(param)
	resp, err := b.blockchainAPI(CreateAndSendVoteDCBBoardTransaction, param)
	if err != nil {
		return errors.Wrap(err, "b.blockchainAPI")
	}

	var (
		data       = resp.(map[string]interface{})
		resultResp = data["Result"]
		errorResp  = data["Error"]
	)
	if errorResp != nil {
		return errors.Errorf("couldn't get result from response data: %+v", errorResp)
	}
	if resultResp == nil {
		return errors.Errorf("couldn't get result from response: req: %+v, data: %+v", param, data)
	}
	return nil
}
