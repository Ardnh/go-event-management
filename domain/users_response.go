package domain

func ToUserResponse(data *User) UserResponse {

	return UserResponse{
		Id:        data.Id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		UserName:  data.Username,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

}
