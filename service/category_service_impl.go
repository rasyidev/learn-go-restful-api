package service

import (
	"context"
	"database/sql"
	"fmt"
	"learn-go-restful-api/helper"
	"learn-go-restful-api/model/domain"
	"learn-go-restful-api/model/web"
	"learn-go-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRespository repository.CategoryRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// sebelum transaksi dimulai, validate dulu requestnya
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// commit atau rollback transaction
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}
	fmt.Printf("CATEGORY (domain.Category): \n %#v", category)

	category = service.CategoryRespository.Save(ctx, tx, category)
	fmt.Printf("CATEGORY (service.CategoryRepository.Save()): \n %#v", category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// sebelum transaksi dimulai, validate dulu requestnya
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// cari data yang mau diupdate dulu
	category, err := service.CategoryRespository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// jika data ada, update nilainya sesuai request
	category.Name = request.Name

	category = service.CategoryRespository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// cari data yang mau dihapus dulu
	category, err := service.CategoryRespository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRespository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRespository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRespository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}
