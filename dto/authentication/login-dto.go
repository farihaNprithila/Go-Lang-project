package authentication

type LoginDto struct {
	UserName string `json:"username" form:"username" binding:"required" `
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}
