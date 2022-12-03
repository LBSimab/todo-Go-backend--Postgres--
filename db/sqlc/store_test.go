package db

import (
	"context"
	"database/sql"
	"testing"
	"todo/util"

	"github.com/stretchr/testify/require"
)

func TestTaskTx(t *testing.T) {

	store := NewStore(testDB)

	errz := make(chan error)
	resultz := make(chan TaskTxResults)

	n := 5
	for i := 0; i < n; i++ {
		account1 := CreateRandomAccount(t)
		go func() {

			result, err := store.TaskTx(context.Background(), TaskTxParams{
				Name: sql.NullString{String: util.RandomName(), Valid: true},

				Supervisor: sql.NullInt32{Int32: int32(account1.ID), Valid: true},
				Category:   sql.NullString{String: util.RandomCategory(), Valid: true},
			})

			errz <- err
			resultz <- result

		}()

	}
	for k := 0; k < n; k++ {
		err := <-errz
		require.NoError(t, err)
		result := <-resultz
		require.NotEmpty(t, result)
		println(result.Task.ID)
		println(result.Task.Supervisor.Int32)

	}
}
