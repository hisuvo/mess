package users

import "messify/internal/domain/users/dto"

func ToUserResponse(user *User) *dto.UserResponse {
	if user == nil {
		return nil
	}
	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		ImageURL:  user.ImageURL,
		Address:   user.Address,
		Role:      string(user.Role),
		Status:    string(user.Status),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []*User) []*dto.UserResponse {
	if users == nil {
		return nil
	}
	var userResponses []*dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}