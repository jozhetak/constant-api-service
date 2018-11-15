package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/ninjadotorg/constant-api-service/dao"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	tokenLength                      = 10
	letters                          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	verificationTokenExpiredDuration = 24 * time.Hour
)

type User struct {
	r  *dao.User
	bc *blockchain.Blockchain
}

func NewUserService(r *dao.User, bc *blockchain.Blockchain) *User {
	return &User{
		r:  r,
		bc: bc,
	}
}

func (u *User) Register(firstName, lastName, email, password, confirmPassword, uType string, publicKey *string) error {
	if email == "" {
		return ErrInvalidEmail
	}
	if password == "" || confirmPassword == "" {
		return ErrInvalidPassword
	}
	if password != confirmPassword {
		return ErrPasswordMismatch
	}
	userType := models.GetUserType(uType)
	if userType == models.InvalidUserType {
		return ErrInvalidUserType
	}

	user, err := u.r.FindByEmail(email)
	if err != nil {
		return errors.Wrap(err, "u.r.FindByEmail")
	}
	if user != nil {
		return ErrEmailAlreadyExists
	}

	paymentAddress, readonlyKey, privKey, err := u.bc.GetAccountWallet(email)
	if err != nil {
		return errors.Wrap(err, "u.bc.GetAccountWallet")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}

	token := u.generateVerificationToken()
	encrypted, err := u.bc.EncryptData(token)
	if err != nil {
		return errors.Wrap(err, "u.bc.EncryptData")
	}
	fmt.Printf("encrypted = %+v\n", encrypted)
	// ul := &models.UserLenderVerification{
	//         UserVerification: models.UserVerification{
	//                 User:      user,
	//                 Token:     token,
	//                 IsValid:   true,
	//                 ExpiredAt: time.Now().Add(verificationTokenExpiredDuration),
	//         },
	//         EncryptedString: encrypted.(string),
	// }
	if err := u.r.Create(&models.User{
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Password:       string(hashed),
		PaymentAddress: paymentAddress,
		ReadonlyKey:    readonlyKey,
		PrivKey:        privKey,
		Type:           userType,
	}); err != nil {
		return errors.Wrap(err, "u.r.Create")
	}

	return nil
}

func (u *User) FindByID(id int) (*models.User, error) {
	user, err := u.r.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "u.r.FindByID")
	}
	return user, nil
}

func (u *User) Authenticate(email, password string) (*UserResp, error) {
	user, err := u.r.FindByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "u.r.FindByEmail")
	}
	if user == nil {
		return nil, ErrEmailNotExists
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidPassword
	}
	return assembleUser(user), nil
}

func (u *User) ForgotPassword(email string) error {
	user, err := u.r.FindByEmail(email)
	if err != nil {
		return errors.Wrap(err, "u.r.FindByEmail")
	}
	if user == nil {
		return ErrInvalidEmail
	}

	if err := u.r.CreateVerification(&models.UserVerification{
		User:      user,
		Token:     u.generateVerificationToken(),
		IsValid:   true,
		ExpiredAt: time.Now().Add(verificationTokenExpiredDuration),
	}); err != nil {
		return errors.Wrap(err, "u.r.CreateRecovery")
	}

	return nil
}

func (u *User) generateVerificationToken() string {
	b := make([]byte, tokenLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (u *User) ResetPassword(token, password, confirmPassword string) error {
	if password == "" || confirmPassword == "" {
		return ErrInvalidPassword
	}
	if password != confirmPassword {
		return ErrPasswordMismatch
	}

	recovery, err := u.r.FindRecoveryToken(token)
	if err != nil {
		return errors.Wrap(err, "u.r.FindRecoveryToken")
	}
	if recovery == nil {
		return ErrInvalidRecoveryToken
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}

	recovery.IsValid = false
	recovery.User.Password = string(hashed)
	if err := u.r.ResetPasswordAndInvalidateToken(recovery); err != nil {
		return errors.Wrap(err, "u.r.Update")
	}

	return nil
}

type UserResp struct {
	ID             uint   `json:"id"`
	UserName       string `json:"user_name"`
	FirstName      string `json:"first_name"`
	LasstName      string `json:"lasst_name"`
	Email          string `json:"email"`
	PaymentAddress string `json:"payment_address"`
	Type           string `json:"type"`
}

func assembleUser(u *models.User) *UserResp {
	return &UserResp{
		ID:             u.ID,
		UserName:       u.UserName,
		FirstName:      u.FirstName,
		LasstName:      u.LastName,
		Email:          u.Email,
		PaymentAddress: u.PaymentAddress,
		Type:           u.Type.String(),
	}
}
