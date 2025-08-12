package helper

import (
	"github.com/syrlramadhan/e-commerce-expo-go-api/dto"
	"github.com/syrlramadhan/e-commerce-expo-go-api/model"
)

func ConvertUserToResponseDTO(user model.User) dto.UserResponse {
	return dto.UserResponse{
		IdUser:  user.IdUser,
		Nama:    user.Nama,
		Email:   user.Email,
		Password: user.Password,
	}
}

func ConvertUserToListResponseDTO(user []model.User) []dto.UserResponse {
	var userResponse []dto.UserResponse
	for _, users := range user {
		userResponse = append(userResponse, ConvertUserToResponseDTO(users))
	}

	return userResponse
}