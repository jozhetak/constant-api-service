package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *Server) Authenticate(c *gin.Context) (interface{}, error) {
	type request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	user, err := s.userSvc.Authenticate(req.Email, req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "u.svc.Authenticate")
	}
	return user, nil
}

func (s *Server) Register(c *gin.Context) {
	type request struct {
		FirstName       string  `json:"first_name"`
		LastName        string  `json:"last_name"`
		Email           string  `json:"email"`
		Password        string  `json:"password"`
		ConfirmPassword string  `json:"confirm_password"`
		Type            string  `json:"type"`
		PublicKey       *string `json:"public_key"`
	}
	type response struct {
		Message string         `json:"Message"`
		Error   *service.Error `json:"Error"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response{Error: service.ErrInvalidArgument})
		return
	}

	err := s.userSvc.Register(req.FirstName, req.LastName, req.Email, req.Password, req.ConfirmPassword, req.Type, req.PublicKey)
	switch err {
	case service.ErrInvalidEmail, service.ErrInvalidPassword, service.ErrInvalidUserType, service.ErrPasswordMismatch, service.ErrEmailAlreadyExists:
		c.JSON(http.StatusBadRequest, response{Error: err.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, response{Message: "Register successfully"})
	default:
		s.logger.Error("u.svc.Register", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
	}
}

func (s *Server) ForgotPassword(c *gin.Context) {
	type request struct {
		Email string `json:"email"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response{Error: service.ErrInvalidArgument})
		return
	}

	err := s.userSvc.ForgotPassword(req.Email)
	switch err := errors.Cause(err); err {
	case service.ErrInvalidEmail:
		c.JSON(http.StatusBadRequest, response{Error: err.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, gin.H{"message": "request successfully"})
	default:
		s.logger.Error("s.userSvc.ForgotPassword", zap.String("email", req.Email), zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
	}
}

func (s *Server) ResetPassword(c *gin.Context) {
	type request struct {
		Token              string `json:"token" binding:"required"`
		NewPassword        string `json:"new_password" binding:"required"`
		ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response{Error: service.ErrInvalidArgument})
		return
	}

	err := s.userSvc.ResetPassword(req.Token, req.NewPassword, req.ConfirmNewPassword)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidPassword, service.ErrPasswordMismatch, service.ErrInvalidRecoveryToken:
		c.JSON(http.StatusBadRequest, response{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, gin.H{"message": "update password successfully"})
	default:
		s.logger.Error("s.userSvc.ResetPassword", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
	}
}

func (s *Server) userFromContext(c *gin.Context) (*models.User, error) {
	userIDVal, ok := c.Get(userIDKey)
	if !ok {
		return nil, errors.New("failed to get userIDKey from context")
	}

	userID := userIDVal.(float64)
	user, err := s.userSvc.FindByID(int(userID))
	if err != nil {
		return nil, errors.Wrap(err, "s.userSvc.FindByID")
	}
	return user, nil
}
