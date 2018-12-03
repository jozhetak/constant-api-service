package serializers

import "github.com/ninjadotorg/constant-api-service/models"

type Resp struct {
	Result interface{} `json:"Result"`
	Error  interface{} `json:"Error"`
}

type UserResp struct {
	ID             uint            `json:"ID"`
	UserName       string          `json:"UserName"`
	FirstName      string          `json:"FirstName"`
	LastName       string          `json:"LastName"`
	Email          string          `json:"Email"`
	PaymentAddress string          `json:"PaymentAddress"`
	Bio            string          `json:"Bio"`
	WalletBalances *WalletBalances `json:"WalletBalances"`
}

func NewUserResp(data models.User) *UserResp {
	result := UserResp{
		PaymentAddress: data.PaymentAddress,
		ID:             data.ID,
		Email:          data.Email,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		UserName:       data.UserName,
		Bio:            data.Bio,
	}
	return &result
}

type UserRegisterResp struct {
	Message         string `json:"Message,omitempty"`
	EncryptedString string `json:"EncryptedString,omitempty"`
}

type UserLoginResp struct {
	Token   string `json:"Token"`
	Expired string `json:"Expired"`
}

type MessageResp struct {
	Message string `json:"Message"`
}
