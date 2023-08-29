package dao

import (
    "context"
	"CwIndustries/dbConfig"
	"CwIndustries/dto"

)

func DB_CreateProjects (application *dto.Projects) error {

	_, err := dbConfig.DATABASE.Collection("Projectss").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}