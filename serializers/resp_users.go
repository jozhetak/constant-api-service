package serializers

type Resp struct {
	Result interface{} `json:"Result"`
	Error  interface{} `json:"Error"`
}

type UserResp struct {
	ID             uint   `json:"ID"`
	UserName       string `json:"UserName"`
	FirstName      string `json:"FirstName"`
	LasstName      string `json:"LasstName"`
	Email          string `json:"Email"`
	PaymentAddress string `json:"PaymentAddress"`
	Type           string `json:"Type"`
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
