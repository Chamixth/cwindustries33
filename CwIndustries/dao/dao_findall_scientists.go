package dao

import (
	"CwIndustries/dbConfig"
	"CwIndustries/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallScientists () (*[]dto.Scientists, error) {
	var objects []dto.Scientists
	results, err := dbConfig.DATABASE.Collection("Scientistss").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Scientists")
	}
	for results.Next(context.Background()) {
		var object dto.Scientists
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Scientists")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
