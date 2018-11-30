package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
)

func (s *Server) Authenticate(c *gin.Context) (interface{}, error) {
	var req serializers.UserReq
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
	var req serializers.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}

	resp, err := s.userSvc.Register(req.FirstName, req.LastName, req.Email, req.Password, req.ConfirmPassword)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidEmail, service.ErrInvalidPassword, service.ErrPasswordMismatch, service.ErrEmailAlreadyExists, service.ErrMissingPubKey:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: resp})
	default:
		s.logger.Error("u.svc.Register", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) ForgotPassword(c *gin.Context) {
	var req serializers.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}

	err := s.userSvc.ForgotPassword(req.Email)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidEmail:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: serializers.MessageResp{Message: "request to reset password successfully"}})
	default:
		s.logger.Error("s.userSvc.ForgotPassword", zap.String("email", req.Email), zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) ResetPassword(c *gin.Context) {
	var req serializers.UserResetPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}

	err := s.userSvc.ResetPassword(req.Token, req.NewPassword, req.ConfirmNewPassword)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidPassword, service.ErrPasswordMismatch, service.ErrInvalidVerificationToken:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: serializers.MessageResp{Message: "update password successfully"}})
	default:
		s.logger.Error("s.userSvc.ResetPassword", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
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

func (s *Server) UpdateUser(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	var req serializers.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}

	if err := s.userSvc.Update(user, &req); err != nil {
		s.logger.Error("s.userSvc.Update", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: "update user successfully"})
}
