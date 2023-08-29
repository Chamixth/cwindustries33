package dao

import (
	"CwIndustries/dbConfig"
	"CwIndustries/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindScientists (objectid string) (*dto.Scientists, error) {
	var object dto.Scientists

	err := dbConfig.DATABASE.Collection("Scientistss").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Scientists")
    }
	if object == (dto.Scientists{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
