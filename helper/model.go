package helper

import (
	"learn-go-restful-api/model/domain"
	"learn-go-restful-api/model/web"
)

// konversi domain.Category ke web.CategoryResponse
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	categoryResponse := []web.CategoryResponse{}
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}

	return categoryResponse
}
