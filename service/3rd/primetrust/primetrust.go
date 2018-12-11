package primetrust

import (
	"encoding/base64"
	"fmt"
)

const (
	Version             = "1.0.0"
	SandboxAPIPrefix    = "https://sandbox.primetrust.com/v2"
	ProductionAPIPrefix = "https://api.primetrust.com/v2"
)

var (
	_apiPrefix  string
	_authHeader string
)

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Init(apiPrefix string, login string, password string) {
	_apiPrefix = apiPrefix
	_authHeader = fmt.Sprintf("Basic %s", basicAuth(login, password))
}
