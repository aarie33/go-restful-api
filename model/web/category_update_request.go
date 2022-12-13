package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required,min=1"`
	Name string `validate:"required,min=3,max=100"`
}
