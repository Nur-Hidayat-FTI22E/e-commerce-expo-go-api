package dto

type UserRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	IdUser   string `json:"id_user"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
