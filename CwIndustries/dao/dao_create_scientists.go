package dao

import (
    "context"
	"CwIndustries/dbConfig"
	"CwIndustries/dto"

)

func DB_CreateScientists (application *dto.Scientists) error {

	_, err := dbConfig.DATABASE.Collection("Scientistss").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}