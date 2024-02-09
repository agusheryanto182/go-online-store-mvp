package entities

import "time"

type User struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	Fullname  string     `gorm:"column:fullname" json:"fullname"`
	Username  string     `gorm:"column:username" json:"username"`
	Email     string     `gorm:"column:email" json:"email"`
	Password  string     `gorm:"column:password" json:"password"`
	Avatar    string     `gorm:"column:avatar" json:"avatar"`
	Phone     *string    `gorm:"column:phone;default:null" json:"phone,omitempty"`
	Address   *string    `gorm:"column:address;default:null" json:"address,omitempty"`
	Role      string     `gorm:"column:role" json:"role"`
	GenderID  *int       `gorm:"column:gender_id;default:null" json:"gender_id,omitempty"`
	Gender    *Gender    `gorm:"foreignKey:GenderID" json:"gender,omitempty"`
	Birthdate *time.Time `gorm:"column:birthdate;default:null" json:"birthdate,omitempty"`
	CreatedAt time.Time  `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type Gender struct {
	ID   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
