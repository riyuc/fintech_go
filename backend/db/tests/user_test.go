package db_test

import (
	"context"
	db "github/riyuc/fintech_backend/db/sqlc"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	arg := db.CreateUserParams{
		Email:          "test@email.com",
		HashedPassword: "test",
	}

	user, err := testQuery.CreateUser(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, arg.HashedPassword, user.HashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2 * time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2 * time.Second)
}