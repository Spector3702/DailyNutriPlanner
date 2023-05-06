package entity

type Foodstuff struct {
	Sno         string `gorm:"primaryKey;size:8" json:"sno"`
	Name        string `gorm:"uniqueIndex" json:"name"`
	OtherName   string `gorm:"column:Other_name" json:"other_name"`
	EnglishName string `gorm:"column:English_name" json:"english_name"`
	Describe    string `json:"describe"`
}

func (Foodstuff) TableName() string {
	return "foodstuff"
}

//`Sno` char(8),
//`Name` varchar(60), -- select ?
//`Other_name` varchar(60),
//`English_name` varchar(200),
//`Describe` varchar(60),
//PRIMARY KEY (`Sno`),
//UNIQUE (`Name`)
