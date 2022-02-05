package foods

import (
	"Daily-Calorie-App-API/business/foods"
	"gorm.io/gorm"
)

type Foods struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Title    string `gorm:"type:varchar(100);not null"`
	ImgURL   string `gorm:"type:varchar(100);not null"`
	Calories float64
	Fat      float64
	Carbs    float64
	Protein  float64
}

func (record *Foods) toDomain() foods.Domain {
	return foods.Domain{
		ID:        int(record.ID),
		Title:     record.Title,
		ImgURL:    record.ImgURL,
		Calories:  record.Calories,
		Fat:       record.Fat,
		Carbs:     record.Carbs,
		Protein:   record.Protein,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain foods.Domain) Foods {
	return Foods{
		ID:       uint(domain.ID),
		Title:    domain.Title,
		ImgURL:   domain.ImgURL,
		Calories: domain.Calories,
		Fat:      domain.Fat,
		Carbs:    domain.Carbs,
		Protein:  domain.Protein,
	}
}

func toDomainArray(record []Foods) []foods.Domain {
	var domains []foods.Domain
	for _, v := range record {
		domains = append(domains, v.toDomain())
	}
	return domains
}
