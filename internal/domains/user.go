package domains

type (
	User struct {
		ID       uint   `gorm:"primary_key"`
		Name     string `gorm:"name"`
		Email    string `gorm:"email"`
		Password string `gorm:"password"`
		Role     string `gorm:"role"`
	}
)
