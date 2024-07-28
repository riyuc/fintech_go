package db_test

import (
	"context"
	db "github/riyuc/fintech_backend/db/sqlc"
	"github/riyuc/fintech_backend/utils"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func CreateRandomUser(t *testing.T) db.User{
	hashedPassword, err := utils.GenerateHashedPassword(utils.RandomString(8))

	if err != nil {
		t.Fatal("Could not generate hashed password", err)	
	}

	arg := db.CreateUserParams{
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := testQuery.CreateUser(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, user.Email, arg.Email)
	assert.Equal(t, user.HashedPassword, arg.HashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2 * time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2 * time.Second)

	user2, err := testQuery.CreateUser(context.Background(), arg)
	assert.Error(t, err)
	assert.Empty(t, user2)

	return user
}

func TestCreateUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	user2, err := testQuery.CreateUser(context.Background(), db.CreateUserParams{
		Email: 		user1.Email,
		HashedPassword: user1.HashedPassword,
	})

	assert.Error(t, err)
	assert.Empty(t, user2)
}

func TestUpdateUser(t *testing.T){
	user := CreateRandomUser(t)

	newPassword, err := utils.GenerateHashedPassword(utils.RandomString(8))

	if err != nil {
		log.Fatal("Could not generate hashed password", err)
	}

	arg := db.UpdateUserPasswordParams{
		HashedPassword: newPassword,
		ID: user.ID,
		UpdatedAt: time.Now(),
	}

	newUser, err := testQuery.UpdateUserPassword(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, newUser.HashedPassword, arg.HashedPassword)
	assert.Equal(t, user.Email, newUser.Email)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2 * time.Second)

}

func TestGetUserById(t *testing.T) {
	user := CreateRandomUser(t)

	user2, err := testQuery.GetUserById(context.Background(), user.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, user2)
	assert.Equal(t, user.Email, user2.Email)
	assert.Equal(t, user.HashedPassword, user2.HashedPassword)
}

func TestGetUserByEmail(t *testing.T) {
	user := CreateRandomUser(t)

	user2, err := testQuery.GetUserByEmail(context.Background(), user.Email)

	assert.NoError(t, err)
	assert.NotEmpty(t, user2)

	assert.Equal(t, user2.HashedPassword, user.HashedPassword)
	assert.Equal(t, user2.Email, user.Email)
}

func TestDeleteUser(t *testing.T) {
	user := CreateRandomUser(t)

	err := testQuery.DeleteUser(context.Background(), user.ID)

	assert.NoError(t, err)

	user2, err := testQuery.GetUserById(context.Background(), user.ID)

	assert.Error(t, err)
	assert.Empty(t, user2)
}