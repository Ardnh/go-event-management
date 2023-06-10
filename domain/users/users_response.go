package domain

func ToUserResponse(data *UserQueryResponse) UserResponse {

	return UserResponse{
		Id:        data.Id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		UserName:  data.Username,
		Email:     data.Email,
		Password:  data.Password,
		RoleId:    data.RoleId,
		Role:      data.Role,
	}

}
