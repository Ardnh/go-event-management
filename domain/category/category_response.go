package domain

func ToCategoryResponse(data *Category) *CategoryResponse {

	return &CategoryResponse{
		Id:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToCategoryResponses(roles []*Category) []*CategoryResponse {
	var response []*CategoryResponse

	for _, role := range roles {
		response = append(response, ToCategoryResponse(role))
	}

	return response
}
