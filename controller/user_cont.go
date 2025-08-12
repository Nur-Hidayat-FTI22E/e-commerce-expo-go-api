package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/e-commerce-expo-go-api/dto"
	"github.com/syrlramadhan/e-commerce-expo-go-api/helper"
	"github.com/syrlramadhan/e-commerce-expo-go-api/service"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	LoginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &userControllerImpl{
		UserService: userService,
	}
}

// CreateUser implements UserController.
func (u *userControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userReq := dto.UserRequest{}
	helper.ReadFromRequestBody(r, &userReq)

	responseDTO, code, err := u.UserService.CreateUser(r.Context(), userReq)
	if err != nil {
		helper.WriteJSONError(w, code, err.Error())
		return
	}

	helper.WriteJSONSuccess(w, responseDTO, "User created successfully")
}

// LoginUser implements UserController.
func (u *userControllerImpl) LoginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	loginReq := dto.LoginRequest{}
	helper.ReadFromRequestBody(r, &loginReq)

	responseDTO, code, err := u.UserService.LoginUser(r.Context(), loginReq.Email, loginReq.Password)
	if err != nil {
		helper.WriteJSONError(w, code, err.Error())
		return
	}

	helper.WriteJSONSuccess(w, responseDTO, "Login successful")
}