package entity

type EverydayNutrition struct {
	Name             string  `gorm:"primaryKey;size:60" json:"name"`
	Analysis_item    string  `gorm:"primaryKey;size:60" json:"analysis_item"`
	Unit             string  `gorm:"size:4" json:"unit"`
	Content_per_unit float64 `gorm:"column:Content_per_unit" json:"content_per_unit"`
}

// Define foreign key relationship with foodstuff
func (e *EverydayNutrition) Foodstuff() Foodstuff {
	return Foodstuff{Name: e.Name}
}

// Define table name for GORM
func (EverydayNutrition) TableName() string {
	return "everyday_nutrition"
}

//`Name` varchar(60), -- select ?
//`Analysis_item` varchar(12),
//`Unit` varchar(4),
//`Content_per_unit` varchar(20),
//`Weight_per_unit` varchar(10),
//`Content_per_unit_weight` float,
//PRIMARY KEY (`Name`,`Analysis_item`),
