package web

// Struct untuk request pada API Create category sesuai API Spec
type CategoryCreateRequest struct {
	Name string `validate:"required,min=2,max=200"`
}
