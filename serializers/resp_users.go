package serializers

type Resp struct {
	Result interface{} `json:"Result"`
	Error  interface{} `json:"Error"`
}

type UserRegisterResp struct {
	EncryptedString *string `json:"EncryptedString"`
}

type UserLoginResp struct {
	Token   string `json:"Token"`
	Expired string `json:"Expired"`
}

type MessageResp struct {
	Message string `json:"Message"`
}
