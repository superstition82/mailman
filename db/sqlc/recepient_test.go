package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateRecepientParams{
		Email:  "chotnt741@gmail.com",
		Status: "unknown",
	}

	recepient, err := testQueries.CreateRecepient(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recepient)

	require.Equal(t, arg.Email, recepient.Email)
	require.Equal(t, arg.Status, recepient.Status)

	require.NotZero(t, recepient.ID)
}
