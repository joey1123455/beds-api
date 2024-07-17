package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joey1123455/beds-api/core/server"
	"github.com/joey1123455/beds-api/internal/common"
	db "github.com/joey1123455/beds-api/internal/db/sqlc"
	"github.com/joey1123455/beds-api/internal/domain"
	"github.com/joey1123455/beds-api/internal/logger"
	"github.com/joey1123455/beds-api/internal/mailer"
	"github.com/joey1123455/beds-api/internal/notifier"
	"github.com/joey1123455/beds-api/internal/security"
	"github.com/joey1123455/beds-api/internal/validator"
)

type Handler struct {
	srv       *server.Server
	Validator security.Validator
}

func NewAuthHandler(srv *server.Server, v security.Validator) *Handler {
	return &Handler{srv: srv, Validator: v}
}

func (auth *Handler) RegisterUser(c *gin.Context) {
	req := domain.RegisterUserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.ErrorLogger(err)
		server.SendBadRequest(c, err)
		return
	}
	v := validator.New()
	if !req.Validate(v) {
		server.SendValidationError(c, validator.NewValidationError(common.ErrValidationFailed, v.Errors))
		return
	}

	// check if user already exists
	if _, err := auth.srv.Store.GetUserByEmail(c, req.Email); err == nil {
		logger.ErrorLogger(fmt.Errorf("while querying user by email-%s during user creation: %v", req.Email, err))
		server.SendBadRequest(c, common.ErrUserUserEmailTaken)
		return
	}

	hashedPassword, err := common.HashPassword(req.Password)
	if err != nil {
		logger.ErrorLogger(err)
		server.SendInternalServerError(c, common.ErrPasswordHashingFailed)
		return
	}

	now := time.Now().UTC()
	expires := now.Add(time.Hour)

	user, err := auth.srv.Store.CreateUser(c, db.CreateUserParams{
		Email:          req.Email,
		Password:       hashedPassword,
		VerifyCode:     db.NewNullString(common.RandomAlphaNumeric(6)),
		UserRole:       db.UserRolesCUSTOMER,
		CodeExpireTime: db.NewNullTime(expires),
		CreatedAt:      now,
		UpdatedAt:      now,
	})
	if err != nil {
		logger.ErrorLogger(err)
		server.SendInternalServerError(c, common.ErrCreatingUser)
		return
	}

	mailData := mailer.EmailData{
		To:      user.Email,
		Subject: "Otp verification Email",
		Data: map[string]interface{}{
			"ExpiryTime": "1 hour",
			"UserID":     user.ID.String(),
			"Code":       user.VerifyCode.String,
		},
	}

	notification := notifier.NewNotification(notifier.NewRecipient(user.Email, ""), "Otp verification Email", common.OtpVerificationEmail, []string{notifier.NotificationChannelEmail})
	notification.AddData("email", mailData)
	if err := auth.srv.TaskDistributor.SendNotification(c.Copy(), notification); err != nil {
		auth.srv.Logger.Error(fmt.Errorf("error sending notification %s", err), nil)
		log.Println(err)
	}

	server.SendCreated(c, "user created", user.ID)

}

func (auth *Handler) ActivateUser(c *gin.Context) {
	req := domain.VerifyOtp{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.ErrorLogger(err)
		server.SendBadRequest(c, common.ErrBadRequest)
		return
	}

	v := validator.New()
	if !req.Validate(v) {
		server.SendValidationError(c, validator.NewValidationError(common.ErrValidationFailed, v.Errors))
		return
	}
	user, err := auth.srv.Store.GetUserByEmail(c, req.Email)
	if err != nil {
		logger.ErrorLogger(err)
		server.SendParsingError(c, common.ErrNotExist)
		return
	}

	if user.EmailVerified {
		server.SendForbidden(c, common.ErrUserEmailAlreadyVerified)
		return
	}
	now := time.Now().UTC()
	if user.CodeExpireTime.Time.Before(now) {
		server.SendError(c, http.StatusGone, common.ErrVerificationCodeExpired)
		return
	}
	if user.VerifyCode.String != req.Code {
		server.SendBadRequest(c, common.ErrInvalidVerificationCode)
		return
	}

	err = auth.srv.Store.ActivateUser(c, db.ActivateUserParams{
		ID:        user.ID,
		UpdatedAt: now,
	})
	if err != nil {
		logger.ErrorLogger(err)
		server.SendParsingError(c, common.ErrFailToProcess)
		return
	}

	server.SendSuccess(c, "user activatiion succesfull", nil)
}

func (auth *Handler) RequestUserActivation(c *gin.Context) {
	email := c.Query("email")
	if !common.ValidateEmail(email) {
		server.SendBadRequest(c, common.ErrInvalidEmail)
		return
	}

	user, err := auth.srv.Store.GetUserByEmail(c, email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.SendParsingError(c, common.ErrNotExist)
			return
		default:
			server.SendParsingError(c, common.ErrFailToProcess)
			return
		}
	}

	if user.EmailVerified {
		server.SendBadRequest(c, common.ErrUserAlreadyVerified)
		return
	}

	now := time.Now().UTC()
	expires := now.Add(time.Hour)
	token, err := auth.srv.Store.ChangeUserVerifyCode(c, db.ChangeUserVerifyCodeParams{
		ID:             user.ID,
		VerifyCode:     db.NewNullString(common.RandomAlphaNumeric(6)),
		CodeExpireTime: db.NewNullTime(expires),
		UpdatedAt:      now,
	})
	if err != nil {
		server.SendParsingError(c, common.ErrFailedToGenerateVerificationCode)
		return
	}

	mailData := mailer.EmailData{
		To:      user.Email,
		Subject: "Otp verification Email",
		Data: map[string]interface{}{
			"ExpiryTime": "1 hour",
			"UserID":     user.ID.String(),
			"Code":       token.String,
		},
	}

	notification := notifier.NewNotification(notifier.NewRecipient(user.Email, ""), "Otp verification Email", common.OtpVerificationEmail, []string{notifier.NotificationChannelEmail})
	notification.AddData("email", mailData)
	if err := auth.srv.TaskDistributor.SendNotification(c.Copy(), notification); err != nil {
		auth.srv.Logger.Error(fmt.Errorf("error sending notification %s", err), nil)
		log.Println(err)
	}

	server.SendSuccess(c, "A verification code has been sent to your email", nil)
}

func (auth *Handler) RequestPasswordReset(c *gin.Context) {
	email := c.Query("email")
	if !common.ValidateEmail(email) {
		server.SendBadRequest(c, common.ErrInvalidEmail)
		return
	}

	user, err := auth.srv.Store.GetUserByEmail(c, email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.SendParsingError(c, common.NewUserNotFound(email))
			return
		default:
			server.SendParsingError(c, common.ErrFailToProcess)
			return
		}
	}

	oldToken, err := auth.srv.Store.GetPasswordResetTokenByUserID(c, user.ID)
	if err == nil {
		err := auth.srv.Store.DeletePasswordResetToken(c, oldToken.Token)
		if err != nil {
			logger.ErrorLogger(err)
		}
	}

	now := time.Now().UTC()
	expires := now.Add(time.Hour)
	token, err := auth.srv.Store.CreateCPasswordResetToken(c, db.CreateCPasswordResetTokenParams{
		UserID:    user.ID,
		Token:     common.RandomAlphaNumeric(6),
		ExpiresAt: expires,
		CreatedAt: db.NewNullTime(now),
	})
	if err != nil {
		server.SendParsingError(c, common.ErrFailedToGeneratePasswordResetCode)
		return
	}

	mailData := mailer.EmailData{
		To:      user.Email,
		Subject: "Password Reset Request Email",
		Data: map[string]interface{}{
			"ExpiryTime": "1 hour",
			"UserID":     user.ID.String(),
			"Code":       token.Token,
		},
	}

	notification := notifier.NewNotification(notifier.NewRecipient(user.Email, ""), "Password Reset verification Email", common.PasswordReset, []string{notifier.NotificationChannelEmail})
	notification.AddData("email", mailData)
	if err := auth.srv.TaskDistributor.SendNotification(c.Copy(), notification); err != nil {
		auth.srv.Logger.Error(fmt.Errorf("error sending notification %s", err), nil)
		log.Println(err)
	}

	server.SendSuccess(c, "A password reset code has been sent to your email", nil)
}

func (auth *Handler) PasswordSetting(c *gin.Context) {
	req := domain.PasswordResetRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		server.SendBadRequest(c, common.ErrBadRequest)
		return
	}

	v := validator.New()
	if !req.Validate(v) {
		server.SendValidationError(c, validator.NewValidationError(common.ErrValidationFailed, v.Errors))
		return
	}

	user, err := auth.srv.Store.GetUserByEmail(c, req.Email)
	if err != nil {
		server.SendParsingError(c, common.NewUserNotFound(req.Email))
		return
	}

	token, err := auth.srv.Store.GetPasswordResetToken(c, req.Code)
	if err != nil {
		server.SendParsingError(c, common.ErrFailedToResetPassword)
	}

	if token.UserID != user.ID {
		server.SendParsingError(c, common.ErrInvalidPasswordResetToken)
		return
	}

	now := time.Now().UTC()
	if token.ExpiresAt.Before(now) {
		err = auth.srv.Store.DeletePasswordResetToken(c, token.Token)
		if err != nil {
			logger.ErrorLogger(err)
		}

		server.SendParsingError(c, common.ErrPasswordResetCodeExpired)
		return
	}

	hashedPassword, err := common.HashPassword(req.Password)
	if err != nil {
		server.SendInternalServerError(c, common.ErrPasswordHashingFailed)
		return
	}

	err = common.CheckPassword(req.Password, user.Password)
	if err == nil {
		server.SendBadRequest(c, common.ErrFailedToResetPassword)
		return
	}

	err = auth.srv.Store.ChangeUserPassword(c, db.ChangeUserPasswordParams{
		ID:        user.ID,
		Password:  hashedPassword,
		UpdatedAt: now,
	})
	if err != nil {
		server.SendParsingError(c, common.ErrFailToProcess)
		return
	}

	err = auth.srv.Store.DeletePasswordResetToken(c, token.Token)
	if err != nil {
		logger.ErrorLogger(err)
	}

	server.SendSuccess(c, "Password saved", user.ID)

}

func (auth *Handler) CreateProfile(c *gin.Context) {
	req := domain.CreateProfileRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		server.SendBadRequest(c, err)
		return
	}
	v := validator.New()
	if !req.Validate(v) {
		server.SendValidationError(c, validator.NewValidationError(common.ErrValidationFailed, v.Errors))
		return
	}

	// check if user already exists
	user, err := auth.srv.Store.GetUserByEmail(c, req.Email)
	if err != nil {
		logger.ErrorLogger(fmt.Errorf("while querying user by email-%s during user creation: %v", req.Email, err))
		server.SendBadRequest(c, common.ErrNotExist)
		return
	}

	now := time.Now().UTC()

	userProfile, err := auth.srv.Store.CreateUserProfile(c, db.CreateUserProfileParams{
		FirstName: sql.NullString{
			String: req.FirstName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: req.LastName,
			Valid:  true,
		},
		Country: sql.NullString{
			String: req.Country,
			Valid:  true,
		},
		PhoneNumber: sql.NullString{
			String: req.PhoneNumber,
			Valid:  true,
		},
		UserID:    user.ID,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		logger.ErrorLogger(err)
		server.SendInternalServerError(c, common.ErrFailToProcess)
		return
	}

	err = auth.srv.Store.UpdateRegistrationStatus(c, db.UpdateRegistrationStatusParams{
		ID: user.ID,
		RegistrationCompleted: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		UpdatedAt: now,
	})
	if err != nil {
		server.SendParsingError(c, common.ErrFailToProcess)
		return
	}

	server.SendCreated(c, "user profile created", userProfile.ID)
}

func (auth *Handler) Login(c *gin.Context) {
	var user db.User
	req := domain.SignInRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		server.SendBadRequest(c, err)
		return
	}

	v := validator.New()
	if !req.Validate(v) {
		server.SendValidationError(c, validator.NewValidationError(common.ErrValidationFailed, v.Errors))
		return
	}

	user, err := auth.srv.Store.GetUserByEmail(c, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			server.SendBadRequest(c, common.ErrIncorrectCredentials)
			return
		}
		server.SendInternalServerError(c, common.ErrCreatingUser)
		return
	}

	err = common.CheckPassword(req.Password, user.Password)
	if err != nil {
		server.SendUnauthorized(c, common.ErrIncorrectCredentials)
		return
	}

	token, err := auth.Validator.GenerateJwtToken(user.ID.String())
	if err != nil {
		server.SendParsingError(c, fmt.Errorf("error generating jwt token"))
	}

	server.SendSuccess(c, "Login Successful", token)
}
