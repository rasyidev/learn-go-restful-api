package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"learn-go-restful-api/helper"
	"learn-go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sqlQuery := "INSERT INTO category(name) values(?)"
	result, err := tx.ExecContext(ctx, sqlQuery, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sqlQuery := "UPDATE category SET name=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sqlQuery, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sqlQuery := "DELETE FROM category WHERE id=?"
	_, err := tx.ExecContext(ctx, sqlQuery, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sqlQuery := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, sqlQuery)
	helper.PanicIfError(err)

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}

func (respository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sqlQuery := "SELECT id, name FROM category WHERE id=?"
	rows, err := tx.QueryContext(ctx, sqlQuery, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil

	} else {
		return category, errors.New(fmt.Sprintf("not found category with id: %v", categoryId))
	}
}
