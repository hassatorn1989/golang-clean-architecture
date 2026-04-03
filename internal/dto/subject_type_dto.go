package dto

type SubjectTypeCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type SubjectTypeUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
