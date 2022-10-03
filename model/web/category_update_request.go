package web

// Struct untuk request pada API Update category sesuai API Spec
type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required, min=1, max=200"`
}
