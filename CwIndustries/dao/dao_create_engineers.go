package dao

import (
    "context"
	"CwIndustries/dbConfig"
	"CwIndustries/dto"

)

func DB_CreateEngineers (application *dto.Engineers) error {

	_, err := dbConfig.DATABASE.Collection("Engineerss").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}