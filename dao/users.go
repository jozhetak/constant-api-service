package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db}
}

func (u *User) Create(user *models.User) error {
	return errors.Wrap(u.db.Create(user).Error, "u.db.Create")
}

func (u *User) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "u.db.Where.First")
	}
	return &user, nil
}

func (u *User) FindByID(id int) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "u.db.Where.First")
	}
	return &user, nil
}

func (u *User) CreateRecovery(r *models.UserRecovery) error {
	return errors.Wrap(u.db.Create(r).Error, "u.db.Create")
}

func (u *User) UpdateRecovery(r *models.UserRecovery) error {
	return errors.Wrap(u.db.Save(r).Error, "u.db.save")
}

func (u *User) FindRecoveryToken(token string) (*models.UserRecovery, error) {
	var r models.UserRecovery
	if err := u.db.Preload("User").Where("token = ?", token).Where("is_valid = 1").First(&r).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "u.db.Where")
	}
	return &r, nil
}

func (u *User) Update(user *models.User) error {
	return errors.Wrap(u.db.Save(user).Error, "u.db.save")
}

func (u *User) ResetPasswordAndInvalidateToken(r *models.UserRecovery) (err error) {
	tx := u.db.Begin()
	if tErr := tx.Error; tErr != nil {
		err = errors.Wrap(tErr, "tx.Error")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = errors.Wrap(tx.Commit().Error, "tx.Commit")
	}()

	if tErr := tx.Save(r.User).Error; tErr != nil {
		err = errors.Wrap(tErr, "tx.Save")
		return
	}
	if tErr := tx.Save(r).Error; tErr != nil {
		err = errors.Wrap(tErr, "tx.Save")
		return
	}
	return
}
