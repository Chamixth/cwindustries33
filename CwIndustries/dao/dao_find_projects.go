package dao

import (
	"CwIndustries/dbConfig"
	"CwIndustries/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindProjects (objectid string) (*dto.Projects, error) {
	var object dto.Projects

	err := dbConfig.DATABASE.Collection("Projectss").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Projects")
    }
	if object == (dto.Projects{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
