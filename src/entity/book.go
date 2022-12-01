package entity

type Book struct {
	ID      uint64 `json:"id" gorm:"primary_key:auto increment"`
	Name    string `json:"name" binding:"required" gorm:"type:varchar(255)"`
	Details string `json:"password" gorm:"type:varchar(255)"`
	UserId  uint64 `json:"-" gorm:"not null"`
	User    User   `json:"user" gorm:"foreignkey:UserID;constrain:onUpdate:CASCADE,onDelete:CASCADE;"`
}
