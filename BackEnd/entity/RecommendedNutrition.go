package entity

// "time"

type RecommendedNutrition struct {
	Gender          string  `gorm:"primary_key;type:enum('male','female')" json:"gender"`
	Age             string  `gorm:"primary_key;type:enum('<1','1-6','7-12','13-15','16-18','19-44','45-64','65-74','75+')" json:"age"`
	NumbersOfPeople int     `gorm:"column:Numbers_of_people" json:"numbers_of_people"`
	Calorie         float64 `json:"calorie"`
	Protein         float64 `json:"protein"`
	Fat             float64 `json:"fat"`
	Carbohydrate    float64 `json:"carbohydrate"`
	VitaminB1       float64 `json:"vitaminB1"`
	VitaminB2       float64 `json:"vitaminB2"`
	VitaminC        float64 `json:"vitaminC"`
	Nicotine        float64 `json:"nicotine"`
	VitaminB6       float64 `json:"vitaminB6"`
	VitaminA        float64 `json:"vitaminA"`
	VitaminE        float64 `json:"vitaminE"`
	Ca              float64 `json:"ca"`
	P               float64 `json:"p"`
	Fe              float64 `json:"fe"`
	Mg              float64 `json:"mg"`
	Zn              float64 `json:"zn"`
	Na              float64 `json:"na"`
	K               float64 `json:"k"`
}
