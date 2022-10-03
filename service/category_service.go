// The bussiness logic

package service

import (
	"context"
	"learn-go-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindAll(ctx context.Context)
	FindById(ctx context.Context)
}
