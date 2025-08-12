package repository

import (
	"context"
	"database/sql"

	"github.com/syrlramadhan/e-commerce-expo-go-api/model"
)

type UserRepository interface {
	AddUser(ctx context.Context, tx *sql.Tx, user model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (model.User, error)
}

type userRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &userRepositoryImpl{}
}

// AddUser implements UserRepository.
func (u *userRepositoryImpl) AddUser(ctx context.Context, tx *sql.Tx, user model.User) (model.User, error) {
	query := "INSERT INTO user (id_user, nama, email, password) VALUES (?, ?, ?, ?)"

	_, err := tx.ExecContext(ctx, query, user.IdUser, user.Nama, user.Email, user.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// GetUserByEmail implements UserRepository.
func (u *userRepositoryImpl) GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (model.User, error) {
	query := "SELECT id_user, nama, email, password FROM user WHERE email = ?"
	row := tx.QueryRowContext(ctx, query, email)
	var user model.User
	err := row.Scan(&user.IdUser, &user.Nama, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}