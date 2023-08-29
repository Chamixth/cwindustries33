package dao

import (
    "context"
	"CwIndustries/dbConfig"
	"CwIndustries/dto"

)

func DB_CreateWeapons (application *dto.Weapons) error {

	_, err := dbConfig.DATABASE.Collection("Weaponss").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}