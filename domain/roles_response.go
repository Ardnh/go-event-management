package domain

func ToRolesResponse(data Roles) RolesResponse {

	return RolesResponse{
		Id:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

}

func ToRolesResponses(roles []Roles) []RolesResponse {
	var response []RolesResponse

	for _, role := range roles {
		response = append(response, ToRolesResponse(role))
	}

	return response
}
