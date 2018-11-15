package serializers

type UserReq struct {
	FirstName       string  `json:"FirstName"`
	LastName        string  `json:"LastName"`
	Email           string  `json:"Email" binding:"required"`
	Password        string  `json:"Password"`
	ConfirmPassword string  `json:"ConfirmPassword"`
	Type            string  `json:"Type"`
	PublicKey       *string `json:"PublicKey"`
}

type UserResetPasswordReq struct {
	Token              string `json:"Token" binding:"required"`
	NewPassword        string `json:"NewPassword" binding:"required"`
	ConfirmNewPassword string `json:"ConfirmNewPassword" binding:"required"`
}

type UserLenderVerificationReq struct {
	Token string `json:"Token"`
}
