package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/moviori/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type AuthRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{pool: pool}
}

func (db *AuthRepository) InsertNewUser(ctx context.Context, NewUser *model.HashedLogin) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()

	id := uuid.New()

	execTag, err := tx.Exec(ctx, "INSERT INTO moviori_profile.auth VALUES($1, $2, $3)", id, NewUser.Login, NewUser.Password)
	if err != nil || execTag.RowsAffected() == 0 {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// GetUserByID function returns user information by ID
func (db *AuthRepository) GetUserByID(ctx context.Context) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	id := uuid.New()
	err = tx.QueryRow(ctx, "SELECT id FROM moviori_profile.auth ").Scan(&id)
	if err != nil {
		return fmt.Errorf("QueryRow(): %w", err)
	}

	return nil
}

func (db *AuthRepository) GetUserByLogin(ctx context.Context, login string) (*model.FullUserModel, error) {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return nil, fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()

	user := &model.FullUserModel{}

	err = tx.QueryRow(ctx, "SELECT id, login, password FROM moviori_profile.auth WHERE login=$1", login).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("QueryRow(): %w", err)
	}
	return user, nil
}
