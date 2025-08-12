package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/syrlramadhan/e-commerce-expo-go-api/dto"
	"github.com/syrlramadhan/e-commerce-expo-go-api/helper"
	"github.com/syrlramadhan/e-commerce-expo-go-api/model"
	"github.com/syrlramadhan/e-commerce-expo-go-api/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user dto.UserRequest) (dto.UserResponse, int, error)
	LoginUser(ctx context.Context, email, password string) (string, int, error)
}

type userServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *sql.DB
}

func NewUserServiceImpl(userRepo repository.UserRepository, db *sql.DB) UserService {
	return &userServiceImpl{
		UserRepo: userRepo,
		DB:       db,
	}
}

// CreateUser implements UserService.
func (u *userServiceImpl) CreateUser(ctx context.Context, userReq dto.UserRequest) (dto.UserResponse, int, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return dto.UserResponse{}, http.StatusInternalServerError, err
	}
	defer tx.Commit()

	pass, err := helper.HashPassword(userReq.Password)
	if err != nil {
		return dto.UserResponse{}, http.StatusInternalServerError, err
	}

	user := model.User{
		IdUser:   uuid.New().String(),
		Nama:     userReq.Nama,
		Email:    userReq.Email,
		Password: pass,
	}

	addedUser, err := u.UserRepo.AddUser(ctx, tx, user)
	if err != nil {
		return dto.UserResponse{}, http.StatusInternalServerError, err
	}

	return helper.ConvertUserToResponseDTO(addedUser), http.StatusOK, nil
}

// LoginUser implements UserService.
func (u *userServiceImpl) LoginUser(ctx context.Context, email string, password string) (string, int, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer tx.Commit()

	user, err := u.UserRepo.GetUserByEmail(ctx, tx, email)
	if err != nil {
		return "", http.StatusNotFound, err
	}

	if !helper.VerifyPassword(user.Password, password) {
		return "", http.StatusBadRequest, fmt.Errorf("invalid nra or password: %v", err)
	}

	token, err := helper.GenerateJWT(email, user.Nama)
	if err != nil {
		return "", http.StatusBadRequest, fmt.Errorf("failed to generate token: %v", err)
	}

	return token, http.StatusOK, nil
}
