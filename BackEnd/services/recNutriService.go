package service

import (
	// "BackEnd/entity"
	"BackEnd/entity"
	model "BackEnd/models"
	"reflect"
	"strconv"
)

var RecNutriService = newRecNutriService()

func newRecNutriService() *recNutriService {
	return &recNutriService{m: model.NewRecNutriModel()}
}

type recNutriService struct {
	m model.RecNutriModel
}

// analysis needed nutrition by given record
func (s *recNutriService) AnalysisByRecord(gender, age string, nutritionStr ...string) (*entity.RecommendedNutrition, error) {
	recordNutri := &entity.RecommendedNutrition{}
	nutritionFields := []string{"Calorie", "Protein", "Fat", "Carbohydrate", "VitaminB1", "VitaminB2", "VitaminC", "Nicotine", "VitaminB6", "VitaminA", "VitaminE", "Ca", "P", "Fe", "Mg", "Zn", "Na", "K"}

	for i, field := range nutritionFields {
		reflect.ValueOf(recordNutri).Elem().FieldByName(field).SetFloat(parseFloat(nutritionStr[i]))
	}

	recmdNutri, err := s.m.GetByGenderAndAge(gender, age)
	analysNutri := subtractRecommended(recmdNutri, recordNutri)
	analysNutri.Gender = gender
	analysNutri.Age = age

	return analysNutri, err
}

func parseFloat(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func subtractRecommended(recmd, record *entity.RecommendedNutrition) *entity.RecommendedNutrition {
	return &entity.RecommendedNutrition{
		Calorie:      recmd.Calorie - record.Calorie,
		Protein:      recmd.Protein - record.Protein,
		Fat:          recmd.Fat - record.Fat,
		Carbohydrate: recmd.Carbohydrate - record.Carbohydrate,
		VitaminB1:    recmd.VitaminB1 - record.VitaminB1,
		VitaminB2:    recmd.VitaminB2 - record.VitaminB2,
		VitaminC:     recmd.VitaminC - record.VitaminC,
		Nicotine:     recmd.Nicotine - record.Nicotine,
		VitaminB6:    recmd.VitaminB6 - record.VitaminB6,
		VitaminA:     recmd.VitaminA - record.VitaminA,
		VitaminE:     recmd.VitaminE - record.VitaminE,
		Ca:           recmd.Ca - record.Ca,
		P:            recmd.P - record.P,
		Fe:           recmd.Fe - record.Fe,
		Mg:           recmd.Mg - record.Mg,
		Zn:           recmd.Zn - record.Zn,
		Na:           recmd.Na - record.Na,
		K:            recmd.K - record.K,
	}
}
