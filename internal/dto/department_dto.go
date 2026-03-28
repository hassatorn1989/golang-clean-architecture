package dto

type DepartmentCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type DepartmentUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
