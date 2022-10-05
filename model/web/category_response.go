package web

// Struct untuk response pada API Find, FindAll, Update, dll category sesuai API Spec
type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
