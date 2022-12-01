package dto

type BookCreateDto struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Details string `json:"details" form:"details" binding:"required"`
	UserId  uint64 `json:"user_id,omitempty" form:"user_id,omitempty"  `
}

type BookUpdateDto struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
	Details string `json:"details" form:"details" binding:"required"`
	UserId  uint64 `json:"user_id,omitempty" form:"user_id,omitempty"  `
}
