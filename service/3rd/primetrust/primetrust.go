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

type Primetrust struct {
	Endpoint      string
	Authorization string
	AccountID     string
}

func Init(apiPrefix string, email string, password string, accountID string) *Primetrust {
	authorization := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", email, password))))

	primetrust := &Primetrust{
		Endpoint:      apiPrefix,
		Authorization: authorization,
		AccountID:     accountID,
	}

	return primetrust
}
