package entity

type Foodstuff struct {
	Sno          string `gorm:"primaryKey" json:"sno"`
	Name         string `gorm:"primaryKey" json:"name"`
	Other_name   string `gorm:"" json:"other_name"`
	English_name string `json:"english_name"`
	Describe     string `json:"discribtion"`

	//?

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