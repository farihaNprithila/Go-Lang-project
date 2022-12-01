package entity

type User struct {
	ID        uint64 `json:"id" gorm:"primary_key:auto increment"`
	UserName  string `json:"username" binding:"required" gorm:"type:varchar(30)"`
	Password  string `json:"password" binding:"required" gorm:"->;<-;not null"`
	FirstName string `json:"firstName" gorm:"type:varchar(40)"`
	LastName  string `json:"lastName" gorm:"type:varchar(40)"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(50)"`
	Token     string `json:"token.omitempty" gorm:"-"`
}
