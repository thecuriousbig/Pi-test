package dto

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserRequest struct {
	Username string `json:"username" valid:"required,length(3|20)"`
	Email    string `json:"email" valid:"required,email"`
}

type CreateUserResponse struct {
	ID uint `json:"id"`
}

type GetUserByIDRequest struct {
	ID uint `param:"id" valid:"required"`
}

type UpdateUserRequest struct {
	ID       uint   `param:"id" valid:"required"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type DeleteUserRequest struct {
	ID uint `param:"id" valid:"required"`
}
