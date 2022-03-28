package foods

import (
	"Daily-Calorie-App-API/business/foods"
	"gorm.io/gorm"
)

type Foods struct {
	gorm.Model
	ID          uint    `gorm:"primary_key" json:"ID,omitempty"`
	Title       string  `gorm:"type:varchar(100);not null;unique" json:"title,omitempty"`
	ImgURL      string  `gorm:"type:varchar(255);not null" json:"imgURL,omitempty"`
	Calories    float64 `gorm:"type:decimal(10,2)" json:"calories,omitempty"`
	Fat         float64 `gorm:"type:decimal(10,2)" json:"fat,omitempty"`
	Carbs       float64 `gorm:"type:decimal(10,2)" json:"carbs,omitempty"`
	Protein     float64 `gorm:"type:decimal(10,2)" json:"protein,omitempty"`
	ServingSize float64 `gorm:"type:decimal(10,2)" json:"servingSize,omitempty"`
}

func (record *Foods) toDomain() foods.Domain {
	return foods.Domain{
		ID:          int(record.ID),
		Title:       record.Title,
		ImgURL:      record.ImgURL,
		Calories:    record.Calories,
		Fat:         record.Fat,
		Carbs:       record.Carbs,
		Protein:     record.Protein,
		ServingSize: record.ServingSize,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func fromDomain(domain foods.Domain) Foods {
	return Foods{
		ID:          uint(domain.ID),
		Title:       domain.Title,
		ImgURL:      domain.ImgURL,
		Calories:    domain.Calories,
		Fat:         domain.Fat,
		Carbs:       domain.Carbs,
		Protein:     domain.Protein,
		ServingSize: domain.ServingSize,
	}
}

func toDomainArray(record []Foods) []foods.Domain {
	var domains []foods.Domain
	for _, v := range record {
		domains = append(domains, v.toDomain())
	}
	return domains
}
