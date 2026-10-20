package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email,max=50"`
	Phone    string `json:"phone" validate:"required,max=20"`
	Password string `json:"password" validate:"required,min=8"`
	ImageUrl string `json:"image_url" validate:"omitempty"`
	Address  string `json:"address" validate:"omitempty"`
	Role     string `json:"role" validate:"omitempty,oneof=ADMIN MANAGER MEMBER"`
	Status   string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type ListUserRequest struct {
	Search string `form:"search"`
	Role   string `form:"role"`
	Status string `form:"status"`
	Sort   string `form:"sort"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name" validate:"omitempty,min=2,max=50"`
	Email    *string `json:"email" validate:"omitempty,email,max=50"`
	Phone    *string `json:"phone" validate:"omitempty,max=20"`
	Password *string `json:"password" validate:"omitempty,min=8"`
	ImageUrl *string `json:"image_url" validate:"omitempty"`
	Address  *string `json:"address" validate:"omitempty"`
	Role     *string `json:"role" validate:"omitempty,oneof=ADMIN MANAGER MEMBER"`
	Status   *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}