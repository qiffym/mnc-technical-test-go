package domains

import "time"

type (
	Transaction struct {
		ID          uint      `gorm:"primary_key"`
		Title       string    `gorm:"title"`
		Amount      int       `gorm:"amount"`
		Description string    `gorm:"description"`
		Date        time.Time `gorm:"date"`
	}
)
