package histories

import (
	"time"
)

type Domain struct {
	ID            int
	UserID        int
	FoodID        int
	Water         int
	TotalCalories float64
	TotalFood     int
	Date          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Service interface {
	CreateHistories(histories *Domain) (*Domain, error)
	GetAllHistoryByUserID(userID int) (*[]Domain, error)
}

type Repository interface {
	Insert(histories *Domain) (*Domain, error)
	Update(id int, histories *Domain) (*Domain, error)
	GetHistoriesByUserIDandDate(userID int, date string) (*Domain, error)
	GetAllHistoryByUserID(userID int) (*[]Domain, error)
}
