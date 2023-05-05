package entity

type PersonalInfo struct {
	Email    string  `gorm:"primaryKey" json:"email"`
	Gender   string  `gorm:"type:ENUM('male','female')" json:"gender"`
	Age      string  `gorm:"type:ENUM('<1','1-6','7-12','13-15','16-18','19-44','45-64','65-74','75+')" json:"age"`
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
	WorkLoad string  `gorm:"column:Work_load;type:ENUM('low','medium','high')" json:"work_load"`
}

func (PersonalInfo) TableName() string {
	return "personal_info"
}
