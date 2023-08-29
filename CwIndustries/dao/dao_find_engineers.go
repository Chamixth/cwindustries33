package dao

import (
	"CwIndustries/dbConfig"
	"CwIndustries/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindEngineers (objectid string) (*dto.Engineers, error) {
	var object dto.Engineers

	err := dbConfig.DATABASE.Collection("Engineerss").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Engineers")
    }
	if object == (dto.Engineers{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
