package serializers

type UserReq struct {
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
	UserName        string `json:"UserName"`
	Bio             string `json:"bio"`
}

type UserResetPasswordReq struct {
	Token              string `json:"Token" binding:"required"`
	NewPassword        string `json:"NewPassword" binding:"required"`
	ConfirmNewPassword string `json:"ConfirmNewPassword" binding:"required"`
}

type UserLenderVerificationReq struct {
	Token string `json:"Token"`
}
