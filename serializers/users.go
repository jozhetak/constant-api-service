package serializers

type UserRequest struct {
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
	Type            string  `json:"type"`
	PublicKey       *string `json:"public_key"`
}
