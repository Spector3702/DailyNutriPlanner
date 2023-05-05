package entity

type EverydayNutrition struct {
	Name                    string `gorm:"" json:"name"`
	Analysis_item           string `gorm:"" json:"analysis_item"`
	Unit                    string `json:"unit"`
	Content_per_unit        string `json:"content_per_unit"`
	Weight_per_unit         string `json:"weight_per_unit"`
	Content_per_unit_weight string `json:"content_per_unit_weight"`

	//?

}

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
