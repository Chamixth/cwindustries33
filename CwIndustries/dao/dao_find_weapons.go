package dao

import (
	"CwIndustries/dbConfig"
	"CwIndustries/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindWeapons (objectid string) (*dto.Weapons, error) {
	var object dto.Weapons

	err := dbConfig.DATABASE.Collection("Weaponss").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Weapons")
    }
	if object == (dto.Weapons{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
