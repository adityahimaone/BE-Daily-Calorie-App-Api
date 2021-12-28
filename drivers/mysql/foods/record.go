package foods

import (
	"Daily-Calorie-App-API/businesses/foods"
	"Daily-Calorie-App-API/drivers/mysql/nutritioninfo"
	"gorm.io/gorm"
)

type Foods struct {
	gorm.Model
	ID              uint   `gorm:"primary_key"`
	Title           string `gorm:"type:varchar(100);not null"`
	ImgURL          string `gorm:"type:varchar(100);not null"`
	NutritionInfoID uint
	NutritionInfo   nutritioninfo.NutritionInfo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Calories        float64
}

func (record *Foods) toDomain() foods.Domain {
	return foods.Domain{
		ID:              int(record.ID),
		Title:           record.Title,
		ImgURL:          record.ImgURL,
		NutritionInfoID: int(record.NutritionInfoID),
		Calories:        record.NutritionInfo.Calories,
		Fat:             record.NutritionInfo.Fat,
		Carbs:           record.NutritionInfo.Carbs,
		Protein:         record.NutritionInfo.Protein,
		ServingSize:     record.NutritionInfo.ServingSize,
		CreatedAt:       record.CreatedAt,
		UpdatedAt:       record.UpdatedAt,
	}
}

func fromDomain(domain foods.Domain) Foods {
	return Foods{
		ID:              uint(domain.ID),
		Title:           domain.Title,
		ImgURL:          domain.ImgURL,
		NutritionInfoID: uint(domain.NutritionInfoID),
	}
}

func toDomainArray(record []Foods) []foods.Domain {
	var domains []foods.Domain
	for _, v := range record {
		domains = append(domains, v.toDomain())
	}
	return domains
}
