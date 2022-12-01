package dto

type UserCreateDto struct {
	UserName  string `json:"username" form:"username" binding:"required" `
	Password  string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required" validate:"email"`
}

type UserUpdateDto struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	UserName  string `json:"username" form:"username" binding:"required" `
	Password  string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required" validate:"email"`
}
