package service

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/ninjadotorg/constant-api-service/dao"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
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

func (u *User) validate(firstName, lastName, email, password, confirmPassword string) error {
	if email == "" {
		return ErrInvalidEmail
	}
	if password == "" || confirmPassword == "" {
		return ErrInvalidPassword
	}
	if password != confirmPassword {
		return ErrPasswordMismatch
	}
	return nil
}

func (u *User) createLenderUser(user *models.User, pubKey string) (*serializers.UserRegisterResp, error) {
	token := u.generateVerificationToken()
	encrypted, err := u.bc.EncryptData(pubKey, token)
	if err != nil {
		return nil, errors.Wrap(err, "u.bc.EncryptData")
	}

	v := &models.UserLenderVerification{
		UserVerification: models.UserVerification{
			User:      user,
			Token:     token,
			IsValid:   true,
			ExpiredAt: time.Now().Add(verificationTokenExpiredDuration),
		},
		EncryptedString: encrypted,
	}
	if err := u.r.CreateLenderUser(user, v); err != nil {
		return nil, errors.Wrap(err, "u.createLenderUser")
	}
	return &serializers.UserRegisterResp{
		EncryptedString: encrypted,
	}, nil
}

func (u *User) createNormalUser(user *models.User) (*serializers.UserRegisterResp, error) {
	if err := u.r.Create(user); err != nil {
		return nil, errors.Wrap(err, "u.r.Create")
	}
	return &serializers.UserRegisterResp{Message: "register successfully"}, nil
}

func (u *User) Register(firstName, lastName, email, password, confirmPassword, uType string, pubKey *string) (*serializers.UserRegisterResp, error) {
	if err := u.validate(firstName, lastName, email, password, confirmPassword); err != nil {
		return nil, errors.Wrap(err, "u.validate")
	}

	userType := models.GetUserType(uType)
	if userType == models.InvalidUserType {
		return nil, ErrInvalidUserType
	}

	user, err := u.r.FindByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "u.r.FindByEmail")
	}
	if user != nil {
		return nil, ErrEmailAlreadyExists
	}

	paymentAddress, readonlyKey, privKey, err := u.bc.GetAccountWallet(email)
	if err != nil {
		return nil, errors.Wrap(err, "u.bc.GetAccountWallet")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}

	user = &models.User{
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Password:       string(hashed),
		PaymentAddress: paymentAddress,
		ReadonlyKey:    readonlyKey,
		PrivKey:        privKey,
		Type:           userType,
		IsActive:       true,
	}
	if userType == models.Lender {
		if pubKey == nil {
			return nil, ErrMissingPubKey
		}

		user.IsActive = false
		return u.createLenderUser(user, *pubKey)
	}
	return u.createNormalUser(user)
}

func (u *User) FindByID(id int) (*models.User, error) {
	user, err := u.r.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "u.r.FindByID")
	}
	return user, nil
}

func (u *User) Authenticate(email, password string) (*serializers.UserResp, error) {
	user, err := u.r.FindByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "u.r.FindByEmail")
	}
	if user == nil {
		return nil, ErrEmailNotExists
	}
	if !user.IsActive {
		return nil, ErrInactiveAccount
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

	v, err := u.r.FindVerificationToken(token)
	if err != nil {
		return errors.Wrap(err, "u.r.FindRecoveryToken")
	}
	if v == nil {
		return ErrInvalidVerificationToken
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}

	v.IsValid = false
	v.User.Password = string(hashed)
	if err := u.r.ResetPassword(v); err != nil {
		return errors.Wrap(err, "u.r.Update")
	}
	return nil
}

func (u *User) VerifyLender(token string) error {
	v, err := u.r.FindLenderVerificationToken(token)
	if err != nil {
		return errors.Wrap(err, "u.r.FindRecoveryToken")
	}
	if v == nil {
		return ErrInvalidVerificationToken
	}

	v.IsValid = false
	v.User.IsActive = true
	if err := u.r.VerifyLender(v); err != nil {
		return errors.Wrap(err, "u.r.VerifyLenderUser")
	}
	return nil
}

func assembleUser(u *models.User) *serializers.UserResp {
	return &serializers.UserResp{
		ID:             u.ID,
		UserName:       u.UserName,
		FirstName:      u.FirstName,
		LasstName:      u.LastName,
		Email:          u.Email,
		PaymentAddress: u.PaymentAddress,
		Type:           u.Type.String(),
	}
}
