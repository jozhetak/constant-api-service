package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/ninjadotorg/constant-api-service/dao"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	emailHelper "github.com/ninjadotorg/constant-api-service/templates/email"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	tokenLength                      = 10
	letters                          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	verificationTokenExpiredDuration = 24 * time.Hour

	resetPasswordBaseURL = "https://google.com"
)

type User struct {
	r      *dao.User
	bc     *blockchain.Blockchain
	mailer *emailHelper.Email
}

func NewUserService(r *dao.User, bc *blockchain.Blockchain, mailer *emailHelper.Email) *User {
	return &User{
		r:      r,
		bc:     bc,
		mailer: mailer,
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

func (u *User) createNormalUser(user *models.User) (*serializers.UserRegisterResp, error) {
	if err := u.r.Create(user); err != nil {
		return nil, errors.Wrap(err, "u.portalDao.Create")
	}
	return &serializers.UserRegisterResp{Message: "register successfully"}, nil
}

func (u *User) Register(firstName, lastName, email, password, confirmPassword string) (*serializers.UserRegisterResp, error) {
	if err := u.validate(firstName, lastName, email, password, confirmPassword); err != nil {
		return nil, errors.Wrap(err, "u.validate")
	}

	user, err := u.r.FindByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "u.portalDao.FindByEmail")
	}
	if user != nil {
		return nil, ErrEmailAlreadyExists
	}

	paymentAddress, readonlyKey, privKey, err := u.bc.GetAccountWallet(email)
	if err != nil {
		return nil, errors.Wrap(err, "u.blockchainService.GetAccountWallet")
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
		IsActive:       true,
	}
	return u.createNormalUser(user)
}

func (u *User) FindByID(id int) (*models.User, error) {
	user, err := u.r.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "u.portalDao.FindByID")
	}
	return user, nil
}

func (u *User) Authenticate(email, password string) (*serializers.UserResp, error) {
	user, err := u.r.FindByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "u.portalDao.FindByEmail")
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
		return errors.Wrap(err, "u.portalDao.FindByEmail")
	}
	if user == nil {
		return ErrInvalidEmail
	}

	token := u.generateVerificationToken()
	if err := u.r.CreateVerification(&models.UserVerification{
		User:      user,
		Token:     token,
		IsValid:   true,
		ExpiredAt: time.Now().UTC().Add(verificationTokenExpiredDuration),
	}); err != nil {
		return errors.Wrap(err, "u.portalDao.CreateRecovery")
	}

	if err := u.mailer.SendForgotPasswordEmail(&emailHelper.ForgotPassword{
		Email: user.Email,
		Name:  user.FirstName,
		Link:  fmt.Sprintf("%s/?token=%s", resetPasswordBaseURL, token),
	}); err != nil {
		return errors.Wrap(err, "u.sendForgotPasswordEmail")
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
		return errors.Wrap(err, "u.portalDao.FindRecoveryToken")
	}
	if v == nil {
		return ErrInvalidVerificationToken
	}
	if !v.IsValid {
		return ErrInvalidVerificationToken
	}
	if v.ExpiredAt.Before(time.Now().UTC()) {
		return ErrInvalidVerificationToken
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}

	v.IsValid = false
	v.User.Password = string(hashed)
	if err := u.r.ResetPassword(v); err != nil {
		return errors.Wrap(err, "u.portalDao.Update")
	}
	return nil
}

func (u *User) Update(user *models.User, req *serializers.UserReq) error {
	if req.UserName != "" {
		user.UserName = req.UserName
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if err := u.r.Update(user); err != nil {
		return errors.Wrap(err, "u.r.Update")
	}
	return nil
}

func assembleUser(u *models.User) *serializers.UserResp {
	return &serializers.UserResp{
		ID:             u.ID,
		UserName:       u.UserName,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		PaymentAddress: u.PaymentAddress,
	}
}
