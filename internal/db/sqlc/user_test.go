package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/joey1123455/beds-api/internal/common"
	"github.com/stretchr/testify/assert"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := common.HashPassword(common.RandomString(6))
	assert.NoError(t, err)
	hashedPin, err := common.HashPassword(fmt.Sprintf("%d", common.RandomInt(100000, 999999)))
	assert.NoError(t, err)

	email := common.RandomEmail()
	username := common.RandomString(6)
	emailVerified := false
	verifyCode := fmt.Sprintf("%d", common.RandomInt(100000, 999999))
	codeExpire := time.Now().Add(4 * time.Minute)
	mfaEnabled := false
	createdAt := time.Now()
	updatedAt := time.Now()

	arg := CreateUserParams{
		Username:       NewNullString(username),
		Email:          email,
		Password:       hashedPassword,
		VerifyCode:     NewNullString(verifyCode),
		CodeExpireTime: NewNullTime(codeExpire),
		Pin:            NewNullString(hashedPin),
		UpdatedAt:      updatedAt,
		CreatedAt:      createdAt,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username.String)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, emailVerified, user.EmailVerified)
	assert.Equal(t, verifyCode, user.VerifyCode.String)
	assert.NotEmpty(t, user.CodeExpireTime)
	assert.Equal(t, mfaEnabled, user.MfaEnabled)
	assert.NotEmpty(t, user.UpdatedAt)
	assert.NotEmpty(t, user.CreatedAt)

	return user
}

func Test_CreateUser(t *testing.T) {

	hashedPassword, err := common.HashPassword(common.RandomString(6))
	assert.NoError(t, err)
	hashedPin, err := common.HashPassword(fmt.Sprintf("%d", common.RandomInt(100000, 999999)))
	assert.NoError(t, err)

	email := common.RandomEmail()
	username := common.RandomString(6)
	emailVerified := false
	verifyCode := fmt.Sprintf("%d", common.RandomInt(100000, 999999))
	codeExpire := time.Now().Add(4 * time.Minute)
	mfaEnabled := false
	createdAt := time.Now()
	updatedAt := time.Now()

	arg := CreateUserParams{
		Username: sql.NullString{
			String: username,
			Valid:  true,
		},
		Email:    email,
		Password: hashedPassword,
		VerifyCode: sql.NullString{
			String: verifyCode,
			Valid:  true,
		},
		Pin:            sql.NullString{String: hashedPin, Valid: true},
		CodeExpireTime: sql.NullTime{Time: codeExpire, Valid: true},
		UpdatedAt:      updatedAt,
		CreatedAt:      createdAt,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username.String)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, hashedPin, user.Pin.String)
	assert.Equal(t, emailVerified, user.EmailVerified)
	assert.Equal(t, verifyCode, user.VerifyCode.String)
	assert.NotEmpty(t, user.CodeExpireTime)
	assert.Equal(t, mfaEnabled, user.MfaEnabled)
	assert.NotEmpty(t, user.UpdatedAt)
	assert.NotEmpty(t, user.CreatedAt)
	// t.Log(err)
}

func Test_GetUser(t *testing.T) {
	user := createRandomUser(t)

	fetchedUser, err := testQueries.GetUser(context.Background(), user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user, fetchedUser)
	assert.Equal(t, user.Password, fetchedUser.Password)
	assert.Equal(t, user.Pin, fetchedUser.Pin)
}

func Test_GetUserByEmail(t *testing.T) {
	user := createRandomUser(t)

	fetchedUser, err := testQueries.GetUserByEmail(context.Background(), user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user, fetchedUser)
	assert.Equal(t, user.Password, fetchedUser.Password)
	assert.Equal(t, user.Pin, fetchedUser.Pin)

}

func Test_GetUsers(t *testing.T) {
	queryLimit := 5
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	fetch5Users, err := testQueries.GetUsers(context.Background(), GetUsersParams{
		Limit:  int32(queryLimit),
		Offset: 0,
	})

	fetchLast2Users, err1 := testQueries.GetUsers(context.Background(), GetUsersParams{
		Limit:  int32(queryLimit),
		Offset: 3,
	})

	assert.Nil(t, err, "expected no error, got %v", err)
	assert.Nil(t, err1, "expected no error, got %v", err)
	assert.Equal(t, queryLimit, len(fetch5Users), "expected %d users, got %d", queryLimit, len(fetch5Users))
	assert.Equal(t, 2, len(fetchLast2Users), "ecpected 2 users, got %d", len(fetchLast2Users))
	assert.Equal(t, user1, fetchLast2Users[0])
	assert.Equal(t, user2, fetchLast2Users[1])

}

func Test_DeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	assert.Nil(t, err)
}

func Test_ActivateUser(t *testing.T) {
	user := createRandomUser(t)
	assert.Equal(t, false, user.EmailVerified, "expected users email verification to be false, got %b", user.EmailVerified)
	err := testQueries.ActivateUser(context.Background(), ActivateUserParams{
		ID:        user.ID,
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	user, err = testQueries.GetUser(context.Background(), user.ID)
	assert.Nil(t, err)
	assert.True(t, user.EmailVerified, "expected users email verification to be true, got %b", user.EmailVerified)

}

func Test_UserEnableMfa(t *testing.T) {
	user := createRandomUser(t)
	assert.Equal(t, false, user.MfaEnabled, "expected users email verification to be false, got %b", user.MfaEnabled)
	err := testQueries.UserEnableMfa(context.Background(), UserEnableMfaParams{
		ID:        user.ID,
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	user, err = testQueries.GetUser(context.Background(), user.ID)
	assert.Nil(t, err)
	assert.True(t, user.MfaEnabled, "expected users email verification to be true, got %b", user.MfaEnabled)
}

func Test_ChangeUserPassword(t *testing.T) {
	user := createRandomUser(t)
	newPassword := common.RandomString(6)
	err := testQueries.ChangeUserPassword(context.Background(), ChangeUserPasswordParams{
		ID:        user.ID,
		Password:  newPassword,
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	user, err = testQueries.GetUser(context.Background(), user.ID)
	assert.Nil(t, err)
	assert.Equal(t, newPassword, user.Password)
}

func Test_ChangeUserVerifyCode(t *testing.T) {
	user := createRandomUser(t)
	newVerifyCode := fmt.Sprintf("%d", common.RandomInt(100000, 999999))
	mewCode, err := testQueries.ChangeUserVerifyCode(context.Background(), ChangeUserVerifyCodeParams{
		ID:         user.ID,
		VerifyCode: NewNullString(newVerifyCode),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	user, err = testQueries.GetUser(context.Background(), user.ID)
	assert.Nil(t, err)
	assert.Equal(t, newVerifyCode, user.VerifyCode.String)
	assert.Equal(t, newVerifyCode, mewCode.String)
}

func Test_ChangeUserPin(t *testing.T) {
	user := createRandomUser(t)
	newPin := common.RandomString(6)
	pin, err := testQueries.ChangeUserPin(context.Background(), ChangeUserPinParams{
		ID:        user.ID,
		Pin:       NewNullString(newPin),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	user, err = testQueries.GetUser(context.Background(), user.ID)
	assert.Nil(t, err)
	assert.Equal(t, newPin, user.Pin.String)
	assert.Equal(t, newPin, pin.String)
}
