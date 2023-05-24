package service

import (
	// "BackEnd/entity"
	"BackEnd/entity"
	model "BackEnd/models"
	"errors"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

var RecNutriService = newRecNutriService()

func newRecNutriService() *recNutriService {
	return &recNutriService{
		m: model.NewRecNutriModel(),
		e: model.NewEverydayNutriModel(),
	}
}

type recNutriService struct {
	m model.RecNutriModel
	e model.EverydayNutriModel
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

// calculate calories based on bmi & workload & age & gender
func (s *recNutriService) CalculateCaloryByInfo(weightStr, heightStr, workload, gender, age string) (float64, error) {
	// calcule factor by base calory
	recmdNutri, err := s.m.GetByGenderAndAge(gender, age)
	baseRecmd, err := s.m.GetByGenderAndAge(gender, "19-44")
	factor := recmdNutri.Calorie / baseRecmd.Calorie

	// calculate bmi
	weight, _ := strconv.ParseFloat(weightStr, 64)
	height, _ := strconv.ParseFloat(heightStr, 64)
	height /= 100
	bmi := weight / (height * height)

	// get catgory by bmi & workload
	bmi_catgory := 0
	workload_catgory := 0

	if bmi < 18.5 {
		bmi_catgory = 0
	} else if bmi < 24 {
		bmi_catgory = 1
	} else {
		bmi_catgory = 2
	}

	if workload == "low" {
		workload_catgory = 0
	} else if workload == "medium" {
		workload_catgory = 1
	} else {
		workload_catgory = 2
	}

	// init matrix for bmi_catgory & workload
	matrix := [3][3]float64{
		{35.0, 30.0, 25.0},
		{40.0, 35.0, 30.0},
		{45.0, 40.0, 35.0},
	}

	// calculate calory by matrix
	calory := matrix[workload_catgory][bmi_catgory] * factor * weight

	return calory, err
}

// return a random food from everydayNutri based on given nutri and need
func (s *recNutriService) GetFoodFromNutriTable(nutri, needStr string) (string, error) {
	need, _ := strconv.ParseFloat(needStr, 64)
	if need < 0 {
		return "stop eating", nil
	}

	foodlist, err := s.e.GetAllFoodByNeedNutri(nutri, needStr)

	if len(*foodlist) == 0 {
		return "", errors.New("no food found")
	}

	// compute a random index
	rand.Seed(time.Now().Unix())
	i := rand.Int() % len(*foodlist)

	return (*foodlist)[i].Name, err // return
}
